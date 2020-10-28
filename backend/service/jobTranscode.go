package service

import (
	"github.com/shenhailuanma/ffmpeg-command-generator/ffmpeg"
	"github.com/sirupsen/logrus"
)

func CreateTranscodeJob(request ffmpeg.FFmpegTranscodeRequest) error {
	cmdString, err := ffmpeg.FFmpegTranscode(request)
	if err != nil {
		return err
	}

	logrus.Info("CreateTranscodeJob, cmdString:", cmdString)

	// create job

	// run job command
	return runJobCommand(cmdString)
}
