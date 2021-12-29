package service

import (
	"fmt"
	"github.com/shenhailuanma/miniTranscoder/config"
	"github.com/shenhailuanma/miniTranscoder/models"
	"github.com/shenhailuanma/miniTranscoder/utils"
)

const templateSnapshot = `#!/bin/sh
ffmpeg -i video.mp4 -f image2 out.png
`

func runSnapshotJob(jobID string) (uint32, error) {
	// generate script file
	var scriptFilePath = fmt.Sprintf("%s/%s/snapshot.sh", config.ConfigDataOutputPath, jobID)
	var logFilePath = fmt.Sprintf("%s/%s/snapshot.log", config.ConfigDataOutputPath, jobID)
	var input = fmt.Sprintf("%s/%s/video.mp4", config.ConfigDataOutputPath, jobID)
	var output = fmt.Sprintf("%s/%s/snapshot.png", config.ConfigDataOutputPath, jobID)

	var scriptString = fmt.Sprintf("#!/bin/sh\nffmpeg -i %s -f image2 %s", input, output)
	err := utils.WriteFile(scriptFilePath, scriptString)
	if err != nil {
		return 0, err
	}

	// update job info
	var snapshot = fmt.Sprintf("%s/%s/snapshot.png", config.ConfigVodFolder, jobID)
	UpdateJobInfo(jobID, models.JobUpdateRequest{Snapshot: &snapshot,})

	// run
	code, err := ScriptRunCommon(jobID, scriptFilePath, logFilePath)
	if err != nil {
		return code, err
	}

	return 0, nil
}
