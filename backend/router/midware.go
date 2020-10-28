package router

import "github.com/gin-gonic/gin"

func CORS() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Max-Age", "86400")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		context.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(200)
		} else {
			context.Next()
		}
	}
}
