package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shenhailuanma/miniTranscoder/controllers"
	"net/http"
)

func Run(listenPort string) error {
	r := gin.Default()

	// web ui sttic files
	//r.StaticFS("/static", http.Dir("/Users/xuzhang/gitbase/miniTranscoder/backend/apidoc"))

	r.StaticFile("/index.html", "/Users/xuzhang/gitbase/miniTranscoder/backend/apidoc/index.html")
	r.StaticFile("/", "/Users/xuzhang/gitbase/miniTranscoder/backend/apidoc/index.html")


	r.StaticFS("/css", http.Dir("/Users/xuzhang/gitbase/miniTranscoder/backend/apidoc/css"))
	r.StaticFS("/fonts", http.Dir("/Users/xuzhang/gitbase/miniTranscoder/backend/apidoc/fonts"))
	r.StaticFS("/img", http.Dir("/Users/xuzhang/gitbase/miniTranscoder/backend/apidoc/img"))
	r.StaticFS("/locales", http.Dir("/Users/xuzhang/gitbase/miniTranscoder/backend/apidoc/locales"))
	r.StaticFS("/utils", http.Dir("/Users/xuzhang/gitbase/miniTranscoder/backend/apidoc/utils"))
	r.StaticFS("/vendor", http.Dir("/Users/xuzhang/gitbase/miniTranscoder/backend/apidoc/vendor"))
	r.StaticFile("/api_data.js", "/Users/xuzhang/gitbase/miniTranscoder/backend/apidoc/api_data.js")
	r.StaticFile("/api_data.json", "/Users/xuzhang/gitbase/miniTranscoder/backend/apidoc/api_data.json")
	r.StaticFile("/api_project.js", "/Users/xuzhang/gitbase/miniTranscoder/backend/apidoc/api_project.js")
	r.StaticFile("/api_project.json", "/Users/xuzhang/gitbase/miniTranscoder/backend/apidoc/api_project.json")
	r.StaticFile("/main.js", "/Users/xuzhang/gitbase/miniTranscoder/backend/apidoc/main.js")

	// healthz
	//r.GET("/healthz", controllers.HealthzController)

	// API
	apiGroup := r.Group("/api")
	apiGroup.Use(CORS())
	{
		apiGroup.POST("/file/upload", controllers.FileUploadController)

		/**
		 * @api {GET} /api/jobs 01-GetJobList
		 * @apiName GetJobList
		 * @apiGroup API
		 * @apiDescription Get all jobs
		 */
		apiGroup.GET("/jobs", controllers.GetJobsController)

		/**
		 * @api {POST} /api/jobs 02-CreateTranscodeJob
		 * @apiName CreateJob
		 * @apiGroup API
		 * @apiDescription Create transcode job
		 * @apiParamExample {json} Request-Example:
		{
		    "inputs":[
		        "/video/Wonders_of_Nature.mp4"
		    ],
		    "outputs":[
		        {
		            "output": "/video/Wonders_of_Nature.output.mp4"
		        }
		    ]
		}
		 *
		 */
		apiGroup.POST("/job/transcode", controllers.CreateTranscodeJobController)
	}

	return r.Run(listenPort)
}
