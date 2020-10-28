package router

import "github.com/gin-gonic/gin"

func Run(listenPort string) error {
	r := gin.Default()


	return r.Run(listenPort)
}
