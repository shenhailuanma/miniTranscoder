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
)

func CreateTranscodeJobController(c *gin.Context)  {
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
	for index,_ := range request.Inputs {
		request.Inputs[index] = config.ConfigDataUploadPath + "/" + request.Inputs[index]
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