package config

var ConfigServicePathBase = "/data"
var ConfigDataUploadPath = "/data/upload"  // user upload file store here
var ConfigDataOutputPath = "/data/vod"     // transcode video store here

func InitDirectoryConfig(base string)  {
	ConfigServicePathBase = base
	ConfigDataUploadPath = ConfigServicePathBase + "/upload"
	ConfigDataOutputPath = ConfigServicePathBase + "/vod"
}


/**
File tree definition:

/ -- ${BaseFolder}
	 	/-- upload
			-- ${SourceVideo1}
			-- ${SourceVideo1}
				...
			-- ${SourceVideoN}
        /-- data
			/--${DataTime}
				/-- ${JobNumber}
						-- job.json
						-- ${CustomName}.mp4
						-- ${CustomName}.m3u8
						/-- m3u8
			/--202112
				...
			/--202201
				/-- 1
				/-- 2
   				  ...
				/-- 102
					-- job.json
					-- ${CustomName}.mp4
					-- ${CustomName}.m3u8
					/-- m3u8

*/