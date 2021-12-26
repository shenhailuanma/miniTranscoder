package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenhailuanma/miniTranscoder/models"
	"github.com/shenhailuanma/miniTranscoder/service"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetJobsController(c *gin.Context)  {
	var response = models.ControllerResponse{}
	response.Status = http.StatusOK
	response.Msg = ""

	jobs, err := service.GetJobList()
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Msg = err.Error()
		logrus.Error("GetJobsController, GetJobList, error:", response.Msg)
		c.JSON(response.Status, &response)
		return
	}

	response.Data = jobs

	c.JSON(response.Status, &response)
}

func RemoveJobController(c *gin.Context)  {
	var response = models.ControllerResponse{}
	response.Status = http.StatusOK
	response.Msg = ""

	err := service.RemoveJob(c.Param("id"))
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Msg = err.Error()
		logrus.Error("RemoveJobController, RemoveJob, error:", response.Msg)
		c.JSON(response.Status, &response)
		return
	}

	c.JSON(response.Status, &response)
}