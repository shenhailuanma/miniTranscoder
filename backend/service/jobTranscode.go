package service

import (
	"errors"
	"github.com/shenhailuanma/ffmpeg-command-generator/ffmpeg"
	"github.com/shenhailuanma/miniTranscoder/models"
	"github.com/shenhailuanma/miniTranscoder/runtime"
	"github.com/shenhailuanma/miniTranscoder/utils"
	"github.com/sirupsen/logrus"
)

func CreateTranscodeJob(request ffmpeg.FFmpegTranscodeRequest) (int, error) {
	// todo: params check
	if len(request.Inputs) == 0 {
		return 0, errors.New("no input")
	}
	if len(request.Outputs) == 0 {
		return 0, errors.New("no output")
	}

	// create job
	var job = models.Job{}
	job.Input = request.Inputs[0]
	job.SourceName = utils.PathLastName(job.Input)
	job.SourceSize = utils.FileSize(job.Input)
	job.Progress = 0
	job.Output = request.Outputs[0].Output

	// generate ffmpeg command
	cmdString, err := ffmpeg.FFmpegTranscode(request)
	if err != nil {
		return 0, err
	}

	logrus.Info("CreateTranscodeJob, cmdString:", cmdString)

	// create job
	job.Command = cmdString
	jobID, err := CreateJob(job)
	if err != nil {
		return 0, err
	}

	// push to job queue
	runtime.JobPush(jobID)

	return jobID, nil
}
