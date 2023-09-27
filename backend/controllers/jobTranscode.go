package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shenhailuanma/ffmpeg-command-generator/ffmpeg"
	"github.com/shenhailuanma/miniTranscoder/config"
	"github.com/shenhailuanma/miniTranscoder/models"
	"github.com/shenhailuanma/miniTranscoder/service"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func CreateTranscodeJobController(c *gin.Context) {
	var response = models.ControllerResponse{}
	response.Status = http.StatusOK
	response.Msg = ""

	var request = ffmpeg.FFmpegTranscodeRequest{}
	err := c.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Msg = err.Error()
		logrus.Error("CreateTranscodeJobController, bind data, error:", response.Msg)
		c.JSON(response.Status, &response)
		return
	}

	logrus.Info("CreateTranscodeJobController, request:", request)

	// check
	if len(request.Inputs) == 0 {
		response.Status = http.StatusBadRequest
		response.Msg = "no input file"
		logrus.Error("CreateTranscodeJobController, bind data, error:", response.Msg)
		c.JSON(response.Status, &response)
		return
	}

	// prepare data
	for inputIndex, inputOne := range request.Inputs {
		/*
			- HTTP(S): http://,https://
			- RTMP: rtmp://, rtmpt://, rtmpe://, rtmpte://, rtmps://, rtmpts://
			- RTSP: rtsp://
			- MMS: mmsh://, mmst://, mms://, mmsu://
			- TCP: tcp://
			- UDP: udp://
			- RTP: rtp://, rtps://
			- SMB: smb://
		*/
		if strings.HasPrefix(inputOne, "http://") ||
			strings.HasPrefix(inputOne, "https://") ||
			strings.HasPrefix(inputOne, "rtmp://") ||
			strings.HasPrefix(inputOne, "rtmpt://") ||
			strings.HasPrefix(inputOne, "rtmpe://") ||
			strings.HasPrefix(inputOne, "rtmpte://") ||
			strings.HasPrefix(inputOne, "rtmps://") ||
			strings.HasPrefix(inputOne, "rtmpts://") ||
			strings.HasPrefix(inputOne, "rtsp://") ||
			strings.HasPrefix(inputOne, "mmsh://") ||
			strings.HasPrefix(inputOne, "mmst://") ||
			strings.HasPrefix(inputOne, "mms://") ||
			strings.HasPrefix(inputOne, "mmsu://") ||
			strings.HasPrefix(inputOne, "tcp://") ||
			strings.HasPrefix(inputOne, "udp://") ||
			strings.HasPrefix(inputOne, "rtp://") ||
			strings.HasPrefix(inputOne, "rtps://") ||
			strings.HasPrefix(inputOne, "smb://") ||
			strings.HasPrefix(inputOne, "ftp://") {

			// support stream type
			logrus.Info("CreateTranscodeJobController, input:", inputOne)
		} else {
			request.Inputs[inputIndex] = config.ConfigDataUploadPath + "/" + inputOne
		}
	}

	if len(request.Outputs) == 0 {
		request.Outputs = []ffmpeg.FFmpegTranscodeOutputParams{}
		var output = ffmpeg.FFmpegTranscodeOutputParams{}

		if output.Format == "" {
			output.Format = "mp4"
		}
		request.Outputs = append(request.Outputs, output)
	}
	request.Globals.Overwrite = true

	jobID, err := service.CreateTranscodeJob(request)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Msg = err.Error()
		logrus.Error("CreateTranscodeJobController, CreateTranscodeJob, error:", response.Msg)
		c.JSON(response.Status, &response)
		return
	}

	response.Data = jobID

	c.JSON(response.Status, &response)
}
