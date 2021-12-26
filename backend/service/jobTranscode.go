package service

import (
	"errors"
	"github.com/shenhailuanma/ffmpeg-command-generator/ffmpeg"
	"github.com/shenhailuanma/miniTranscoder/models"
	"github.com/shenhailuanma/miniTranscoder/runtime"
	"github.com/shenhailuanma/miniTranscoder/utils"
	"github.com/sirupsen/logrus"
)

func CreateTranscodeJob(request ffmpeg.FFmpegTranscodeRequest) (string, error) {
	// todo: params check
	if len(request.Inputs) == 0 {
		return "", errors.New("no input")
	}
	if len(request.Outputs) == 0 {
		return "", errors.New("no output")
	}

	// todo: support muti input and output

	// create job
	var job = models.Job{}
	job.Input = request.Inputs[0]
	job.SourceName = utils.PathLastName(job.Input)
	job.SourceSize = utils.FileSize(job.Input)
	job.Progress = 0

	jobID, err := CreateJob(job)
	if err != nil {
		return "", err
	}

	// update jot output path
	// create job
	job.OutputFormat = request.Outputs[0].Format
	job.Output = JobOutputPath(jobID, job.OutputFormat)

	request.Outputs[0].Output = job.Output

	// generate ffmpeg command
	cmdString, err := ffmpeg.FFmpegTranscode(request)
	if err != nil {
		return "", err
	}

	job.Command = cmdString
	logrus.Info("CreateTranscodeJob, cmdString:", cmdString)

	// update cmdString
	UpdateJobCmdAndOutput(jobID, cmdString, job.Output)


	// push to job queue
	runtime.JobPush(jobID)

	return jobID, nil
}
