package service

import (
	"github.com/shenhailuanma/miniTranscoder/config"
	"github.com/shenhailuanma/miniTranscoder/utils"
)

func PrepareServiceRequiredFolders() error {
	var err error = nil


	err = utils.CreatePath(config.ConfigDataUploadPath)
	if err != nil {
		return err
	}
	err = utils.CreatePath(config.ConfigDataOutputPath)
	if err != nil {
		return err
	}

	return err
}