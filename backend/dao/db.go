package dao

import (
	"github.com/shenhailuanma/miniTranscoder/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var gDatabase *gorm.DB = nil

func init() {
	// logrus.Info("init DB")
	_, err := DatabaseOpen()
	if err != nil {
		logrus.Errorf("DatabaseOpen error:%s\n", err.Error())
	}
}

func DatabaseOpen() (*gorm.DB, error) {

	if gDatabase != nil {
		return gDatabase, nil
	}

	var err error

	gDatabase, err = gorm.Open(sqlite.Open(config.ConfigServiceSqlitePath), &gorm.Config{})
	if err != nil {
		gDatabase = nil
		return nil, err
	}

	return gDatabase, err
}