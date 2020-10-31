package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenhailuanma/miniTranscoder/config"
	"github.com/shenhailuanma/miniTranscoder/models"
	"github.com/sirupsen/logrus"
	"net/http"
)

func FileUploadController(c *gin.Context) {
	var response = models.ControllerResponse{}
	response.Status = http.StatusOK
	response.Msg = ""

	file, err := c.FormFile("file")
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Msg = err.Error()
		logrus.Error("FileUploadController, error:", response.Msg)
		c.JSON(response.Status, &response)
		return
	}
	logrus.Info("FileUploadController, file name:", file.Filename)
	logrus.Info("FileUploadController, file size:", file.Size)

	var path = config.ConfigDataUploadPath + "/" + file.Filename
	ok := c.SaveUploadedFile(file, path)
	if ok != nil {
		response.Status = http.StatusBadRequest
		response.Msg = "save upload file failed"
		logrus.Error("FileUploadController, SaveUploadedFile error:", response.Msg)
		c.JSON(response.Status, &response)
		return
	}

	c.JSON(response.Status, &response)
}
