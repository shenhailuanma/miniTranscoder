package config

const ConfigUploadFolder = "upload"
const ConfigVodFolder = "vod"

var ConfigServicePathBase = "/tmp"
var ConfigDataUploadPath = "/tmp/upload"  // user upload file store here
var ConfigDataOutputPath = "/tmp/vod"     // transcode video store here

func InitDirectoryConfig(base string)  {
	ConfigServicePathBase = base
	ConfigDataUploadPath = ConfigServicePathBase + "/" + ConfigUploadFolder
	ConfigDataOutputPath = ConfigServicePathBase + "/" + ConfigVodFolder
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