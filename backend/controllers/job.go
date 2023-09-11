package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shenhailuanma/miniTranscoder/models"
	"github.com/shenhailuanma/miniTranscoder/service"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetJobsController(c *gin.Context) {
	var response = models.ControllerResponse{}
	response.Status = http.StatusOK
	response.Msg = ""

	jobs, err := service.GetAllJobsInfo()
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

func GetUndoneJobsController(c *gin.Context) {
	var response = models.ControllerResponse{}
	response.Status = http.StatusOK
	response.Msg = ""

	jobs, err := service.GetAllJobsInfo()
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Msg = err.Error()
		logrus.Error("GetJobsController, GetJobList, error:", response.Msg)
		c.JSON(response.Status, &response)
		return
	}

	var outputs = []models.Job{}
	for _, jobOne := range jobs {
		if jobOne.Status != models.JobStatusDone && jobOne.Status != models.JobStatusError {
			outputs = append(outputs, jobOne)
		}
	}

	response.Data = outputs

	c.JSON(response.Status, &response)
}

func GetJobInfoController(c *gin.Context) {
	var response = models.ControllerResponse{}
	response.Status = http.StatusOK
	response.Msg = ""

	jobInfo, err := service.GetJob(c.Param("id"))
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Msg = err.Error()
		logrus.Error("GetJobInfoController, GetJob, error:", response.Msg)
		c.JSON(response.Status, &response)
		return
	}

	response.Data = jobInfo

	c.JSON(response.Status, &response)
}

func UpdateJobController(c *gin.Context) {
	var response = models.ControllerResponse{}
	response.Status = http.StatusOK
	response.Msg = ""

	jobID := c.Param("id")

	var request = models.JobUpdateRequest{}
	err := c.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Msg = err.Error()
		logrus.Error("UpdateJobController, bind data, error:", response.Msg)
		c.JSON(response.Status, &response)
		return
	}

	logrus.Info("UpdateJobController, request:", request)

	err = service.UpdateJob(jobID, request)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Msg = err.Error()
		logrus.Error("UpdateJobController, RemoveJob, error:", response.Msg)
		c.JSON(response.Status, &response)
		return
	}

	c.JSON(response.Status, &response)
}

func RemoveJobController(c *gin.Context) {
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
