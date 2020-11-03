package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenhailuanma/miniTranscoder/models"
	"github.com/shenhailuanma/miniTranscoder/service"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func GetJobsController(c *gin.Context)  {
	var response = models.ControllerResponse{}
	response.Status = http.StatusOK
	response.Msg = ""

	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	jobs, err := service.GetJobList(page, size)
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

func GetJobsCountController(c *gin.Context)  {
	var response = models.ControllerResponse{}
	response.Status = http.StatusOK
	response.Msg = ""

	jobs, err := service.GetJobsCount()
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
