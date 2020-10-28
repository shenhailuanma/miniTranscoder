package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthzController(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}