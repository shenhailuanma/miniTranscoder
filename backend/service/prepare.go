package service

import (
	"github.com/shenhailuanma/miniTranscoder/config"
	"github.com/shenhailuanma/miniTranscoder/dao"
	"github.com/shenhailuanma/miniTranscoder/utils"
)

func PrepareServiceRequiredFolders() error {
	var err error = nil

	err = utils.CreatePath(config.ConfigServicePathWeb)
	if err != nil {
		return err
	}
	err = utils.CreatePath(config.ConfigServicePathBin)
	if err != nil {
		return err
	}
	err = utils.CreatePath(config.ConfigDataUploadPath)
	if err != nil {
		return err
	}
	err = utils.CreatePath(config.ConfigDataOutputPath)
	if err != nil {
		return err
	}
	err = utils.CreatePath(config.ConfigServiceSqliteDir)
	if err != nil {
		return err
	}

	return err
}

func PrepareDatabase() error {

	exist := utils.CheckFileExist(config.ConfigServiceSqlitePath)
	if exist == false {
		// create file
		err := utils.WriteFile(config.ConfigServiceSqlitePath,"")
		if err != nil {
			return err
		}

		err = dao.PrepareJobsTable()
		if err != nil {
			return err
		}
	}

	return nil
}