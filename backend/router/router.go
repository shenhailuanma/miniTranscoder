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
	r.StaticFile("/index.html", config.ConfigServicePathBase+"/ui/index.html")
	r.StaticFile("/", config.ConfigServicePathBase+"/ui/index.html")
	r.StaticFS("/static", http.Dir(config.ConfigServicePathBase+"/ui/static"))

	// web ui download file
	r.StaticFS("/"+config.ConfigVodFolder, http.Dir(config.ConfigDataOutputPath))

	// healthz
	//r.GET("/healthz", controllers.HealthzController)

	// API
	apiGroup := r.Group("/api")
	apiGroup.Use(CORS())
	{
		apiGroup.POST("/file/upload", controllers.FileUploadController)

		apiGroup.GET("/playlist", controllers.GetPlaylistController)

		/**
		 * @api {GET} /api/jobs 01-GetJobList
		 * @apiName GetJobList
		 * @apiGroup API
		 * @apiDescription Get job list
		 */
		apiGroup.GET("/jobs", controllers.GetJobsController)
		apiGroup.GET("/jobs/undone", controllers.GetUndoneJobsController)

		apiGroup.GET("/job/:id", controllers.GetJobInfoController)
		apiGroup.DELETE("/job/:id", controllers.RemoveJobController)

		apiGroup.PUT("/job/:id", controllers.UpdateJobController)
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
