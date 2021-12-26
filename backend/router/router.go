package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shenhailuanma/miniTranscoder/config"
	"github.com/shenhailuanma/miniTranscoder/controllers"
	"net/http"
)

func Run(listenPort string) error {
	r := gin.Default()

	// web ui static files
	r.StaticFile("/index.html", "/miniTranscoder/www/dist/index.html")
	r.StaticFile("/", "/miniTranscoder/www/dist/index.html")
	r.StaticFS("/static", http.Dir("/miniTranscoder/www/dist/static"))

	// web ui download file
	r.StaticFS(config.ConfigServicePathBase, http.Dir(config.ConfigServicePathBase))

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

		apiGroup.DELETE("/job/:id", controllers.RemoveJobController)

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
		            "output": "/video/Wonders_of_Nature.output.mp4",
					"format": "mp4", // default
					"streams": [
						{
							"kind": "video",
							"video": {
								"codec": "h264",
								"preset": "slow",
								"fps": 25,
								"width": 1920,
								"height": 1080
							}
						},
						{
							"kind": "audio",
							"audio": {
								"codec": "aac",
								"channles": 2,
								"": 44100
		                    }
						}
					]
		        }
		    ]
		}
		 *
		 */
		apiGroup.POST("/job/transcode", controllers.CreateTranscodeJobController)
	}

	return r.Run(listenPort)
}
