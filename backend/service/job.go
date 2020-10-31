package service

import (
	"github.com/go-cmd/cmd"
	"github.com/shenhailuanma/miniTranscoder/dao"
	"github.com/shenhailuanma/miniTranscoder/models"
	"github.com/shenhailuanma/miniTranscoder/runtime"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func init()  {
	logrus.Info("service.job init")
	go func() {
		loopDoJob()
	}()
	logrus.Info("service.job init done")
}

func GetJobList() ([]models.Job, error) {
	return dao.GetJobs()
}

func CreateJob(job models.Job) (int, error) {
	return dao.CreateJob(job)
}

func loopDoJob()  {
	for {
		time.Sleep(time.Second * 1)

		jobID := runtime.JobPop()
		logrus.Info("loopDoJob, get job:", jobID)

		// get job
		job, err := dao.GetJobInfo(jobID)
		if err != nil {
			logrus.Error("loopDoJob, GetJobInfo error:", err.Error())
			continue
		}

		logrus.Info("loopDoJob, Command:", job.Command)

		_, err = runCommand("ffmpeg", job.Command)
		if err != nil {
			logrus.Error("loopDoJob, GetJobInfo error:", err.Error())
			continue
		}

	}
}

func runCommand(bin, cmdString string) (int, error) {
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

	command := cmd.NewCmdOptions(cmdOptions, bin, args...)

	// Print STDOUT and STDERR lines streaming from Cmd
	//go func() {
	//	for {
	//		select {
	//		case line := <-command.Stdout:
	//			fmt.Println(line)
	//		case line := <-command.Stderr:
	//			fmt.Println(line)
	//		}
	//	}
	//}()

	// Run and wait for Cmd to return Status
	status := <-command.Start()

	// Cmd has finished but wait for goroutine to print all lines
	for len(command.Stdout) > 0 || len(command.Stderr) > 0 {
		time.Sleep(10 * time.Millisecond)
	}

	return status.Exit,nil
}