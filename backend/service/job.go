package service

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/go-cmd/cmd"
	"github.com/shenhailuanma/miniTranscoder/cache"
	"github.com/shenhailuanma/miniTranscoder/config"
	"github.com/shenhailuanma/miniTranscoder/models"
	"github.com/shenhailuanma/miniTranscoder/runtime"
	"github.com/shenhailuanma/miniTranscoder/utils"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func init() {
	logrus.Info("service.job init")
	go func() {
		loopDoJob()
	}()
	logrus.Info("service.job init done")
}

func CreateJob(job models.Job) (string, error) {
	return InitJob(job)
}

func GetJob(jobID string) (models.Job, error) {
	return GetJobConfig(jobID)
}

func UpdateJob(jobID string, request models.JobUpdateRequest) error {
	return UpdateJobInfo(jobID, request)
}

func RemoveJob(jobID string) error {
	logrus.Info("RemoveJob, clean cache")
	cache.CleanCacheJobs()

	logrus.Info("RemoveJob, jobID:", jobID)
	jobFolder := fmt.Sprintf("%s/%s", config.ConfigDataOutputPath, jobID)
	return utils.RemoveDir(jobFolder)
}

func loopDoJob() {
	for {
		time.Sleep(time.Second * 1)

		jobID := runtime.JobPop()
		logrus.Info("loopDoJob, get job:", jobID)

		// get job
		job, err := GetJobConfig(jobID)
		if err != nil {
			logrus.Error("loopDoJob, GetJobConfig error:", err.Error())
			continue
		}

		logrus.Info("loopDoJob, Command:", job.Command)

		// update status
		UpdateJobStatus(jobID, models.JobStatusProgressing)

		existCode, err := runJob(jobID, job.Command)
		if err != nil {
			logrus.Error("loopDoJob, GetJobInfo error:", err.Error())
			UpdateJobStatus(jobID, models.JobStatusError)
			continue
		}
		logrus.Info("loopDoJob, job:", jobID, " over, exist code:", existCode)

		time.Sleep(time.Second * 1)

		// update progress
		UpdateJobProgress(jobID, 100)

		// update status
		UpdateJobStatus(jobID, models.JobStatusDone)

		// update output file size
		SyncJobOutputFileSize(jobID)

		// snapshot
		runSnapshotJob(jobID)
	}
}

func runCommand(bin, cmdString string, jobID int) (int, error) {
	var args = []string{}
	params := strings.Split(cmdString, " ")

	for _, paramOne := range params {
		if paramOne != "" {
			args = append(args, paramOne)
		}
	}

	// Disable output buffering, enable streaming
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}

	jobCommand := cmd.NewCmdOptions(cmdOptions, bin, args...)
	jobCommand.Env = os.Environ()

	// Print STDOUT and STDERR lines streaming from Cmd
	go func() {
		for {
			select {
			case line := <-jobCommand.Stdout:
				fmt.Println("stdout:" + line)
				jobStatus := jobCommand.Status()
				if jobStatus.Complete {
					logrus.Info("job completed, go routine over")
					return
				}
			case line := <-jobCommand.Stderr:
				fmt.Println("stderr:" + line)
				jobStatus := jobCommand.Status()
				if jobStatus.Complete {
					logrus.Info("job completed, go routine over")
					return
				}
			}
		}
	}()

	// Run and wait for Cmd to return Status
	status := <-jobCommand.Start()

	// waiting for goroutine to print all lines
	time.Sleep(200 * time.Millisecond)

	// Cmd has finished but wait for goroutine to print all lines
	for len(jobCommand.Stdout) > 0 || len(jobCommand.Stderr) > 0 {
		time.Sleep(10 * time.Millisecond)
	}

	return status.Exit, nil
}

func runJob(jobID string, cmdString string) (uint32, error) {
	// generate script file
	var scriptFilePath = fmt.Sprintf("%s/%s/job.sh", config.ConfigDataOutputPath, jobID)
	var logFilePath = fmt.Sprintf("%s/%s/job.log", config.ConfigDataOutputPath, jobID)

	var scriptString = fmt.Sprintf("#!/bin/sh\nffmpeg %s", cmdString)
	err := utils.WriteFile(scriptFilePath, scriptString)
	if err != nil {
		return 0, err
	}

	// run
	return ScriptRun(jobID, scriptFilePath, logFilePath)
}

func ScriptRun(jobID string, scriptPath string, logFilePath string) (uint32, error) {
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return 0, err
	}

	exist, _ := utils.PathExists(scriptPath)
	if exist == false {
		logrus.Info("ScriptRun, script:", scriptPath, ", not exist.")
		return 1, errors.New("script not exist.")
	}

	logrus.Info("ScriptRun, run script:", scriptPath, " begin")

	cmd := exec.Command("sh", scriptPath)
	cmd.Env = os.Environ()

	cmd.Stdout = logFile
	cmd.Stderr = logFile

	// get job progress
	var scriptRunOver = false
	defer func() {
		scriptRunOver = true
	}()

	go func() {
		for {
			time.Sleep(time.Second * 1)
			if scriptRunOver {
				logrus.Info("ScriptRun, script run over, return")
				return
			}
			// update progress
			UpdateJobProgress(jobID, parseFFmpegLogProgress(logFilePath))

			// update output file size
			SyncJobOutputFileSize(jobID)
		}
	}()

	// run
	err = cmd.Run()
	if err != nil {
		logrus.Error("ScriptRun, run script:", scriptPath, ", error:", err)

		// exit code
		code, ok := cmd.ProcessState.Sys().(syscall.WaitStatus)
		if ok {
			return uint32(code.ExitStatus()), err
		}

		return 1, err
	}

	logrus.Info("ScriptRun, run script:", scriptPath, " over")

	return 0, nil
}

func ScriptRunCommon(jobID string, scriptPath string, logFilePath string) (uint32, error) {
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return 0, err
	}

	exist, _ := utils.PathExists(scriptPath)
	if exist == false {
		logrus.Info("ScriptRun, script:", scriptPath, ", not exist.")
		return 1, errors.New("script not exist.")
	}

	logrus.Info("ScriptRun, run script:", scriptPath, " begin")

	cmd := exec.Command("sh", scriptPath)
	cmd.Env = os.Environ()

	cmd.Stdout = logFile
	cmd.Stderr = logFile

	// get job progress
	var scriptRunOver = false
	defer func() {
		scriptRunOver = true
	}()

	go func() {
		for {
			time.Sleep(time.Second * 1)
			if scriptRunOver {
				logrus.Info("ScriptRun, script run over, return")
				return
			}
		}
	}()

	// run
	err = cmd.Run()
	if err != nil {
		logrus.Error("ScriptRun, run script:", scriptPath, ", error:", err)

		// exit code
		code, ok := cmd.ProcessState.Sys().(syscall.WaitStatus)
		if ok {
			return uint32(code.ExitStatus()), err
		}

		return 1, err
	}

	logrus.Info("ScriptRun, run script:", scriptPath, " over")

	return 0, nil
}

func parseFFmpegLogProgress(logPath string) int {
	var progress = 0
	var fileDuration = 0
	var currentDuration = 0

	f, err := os.Open(logPath)
	if err != nil {
		return progress
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\r')
		if err != nil {
			break
		}

		if fileDuration == 0 {
			// get duration
			fileDuration = parseFileDuration(line)
		} else {
			// get latest duration
			currentDuration = parseCurrentDuration(line)
		}
	}

	if currentDuration != 0 && fileDuration != 0 {
		progress = (currentDuration * 100) / fileDuration
	}

	logrus.Info("parseFFmpegLogProgress, fileDuration:", fileDuration, ", currentDuration:", currentDuration, ", progress", progress)

	return progress
}

func parseFileDuration(line string) int {
	var duration = 0

	var re = regexp.MustCompile(`Duration:(.*):(.*):(.*)\.(.*), start`)

	groups := re.FindStringSubmatch(line)

	if len(groups) > 4 {
		hour, err := strconv.Atoi(strings.TrimSpace(groups[1]))
		if err != nil {
			logrus.Error("parse hour error:", err.Error())
			return duration
		}
		minutes, err := strconv.Atoi(groups[2])
		if err != nil {
			logrus.Error("parse minutes error:", err.Error())
			return duration
		}
		seconds, err := strconv.Atoi(groups[3])
		if err != nil {
			logrus.Error("parse seconds error:", err.Error())
			return duration
		}

		duration = 3600*hour + 60*minutes + seconds
	}

	return duration
}

func parseCurrentDuration(line string) int {
	var duration = 0

	var re = regexp.MustCompile(`frame=(.*)time=(.*):(.*):(.*)\.(.*)bitrate=`)

	groups := re.FindStringSubmatch(line)

	if len(groups) > 5 {
		hour, err := strconv.Atoi(groups[2])
		if err != nil {
			logrus.Error("parse hour error:", err.Error())
			return duration
		}
		minutes, err := strconv.Atoi(groups[3])
		if err != nil {
			logrus.Error("parse minutes error:", err.Error())
			return duration
		}
		seconds, err := strconv.Atoi(groups[4])
		if err != nil {
			logrus.Error("parse seconds error:", err.Error())
			return duration
		}

		duration = 3600*hour + 60*minutes + seconds
	}

	return duration
}
