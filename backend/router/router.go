package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shenhailuanma/miniTranscoder/controllers"
)

func Run(listenPort string) error {
	r := gin.Default()

	// healthz
	r.GET("/healthz", controllers.HealthzController)

	// API
	apiGroup := r.Group("/api")
	apiGroup.Use(CORS())
	{
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
