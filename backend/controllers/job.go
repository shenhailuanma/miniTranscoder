package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenhailuanma/miniTranscoder/models"
	"net/http"
)

func GetJobsController(c *gin.Context)  {
	var response = models.ControllerResponse{}
	response.Status = http.StatusOK
	response.Msg = ""


	c.JSON(response.Status, &response)
}
