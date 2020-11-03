package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shenhailuanma/miniTranscoder/controllers"
	"net/http"
)

func Run(listenPort string) error {
	r := gin.Default()

	// web ui static files
	r.StaticFile("/index.html", "/Users/xuzhang/gitbase/miniTranscoder/frontend/dist/index.html")
	r.StaticFile("/", "/Users/xuzhang/gitbase/miniTranscoder/frontend/dist/index.html")
	r.StaticFS("/static", http.Dir("/Users/xuzhang/gitbase/miniTranscoder/frontend/dist/static"))

	// web ui download file
	r.StaticFS("/miniTranscoder", http.Dir("/Users/xuzhang/gitbase/miniTranscoder/backend/miniTranscoder"))

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

		apiGroup.GET("/jobs/count", controllers.GetJobsCountController)

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
