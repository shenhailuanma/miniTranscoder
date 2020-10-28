package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shenhailuanma/ffmpeg-command-generator/ffmpeg"
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

	err = service.CreateTranscodeJob(request)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Msg = err.Error()
		logrus.Error("CreateTranscodeJobController, CreateTranscodeJob, error:", response.Msg)
		c.JSON(response.Status, &response)
		return
	}

	c.JSON(response.Status, &response)
}