package main

import (
	"flag"
	"fmt"
	"github.com/shenhailuanma/miniTranscoder/config"
	"github.com/shenhailuanma/miniTranscoder/router"
	"github.com/shenhailuanma/miniTranscoder/service"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	port = flag.Int("p", 9000, "custom service port")
	directory = flag.String("d", "/tmp", "video data store base directory")
)

func main() {

	flag.Parse()

	logrus.Info("Service start")

	config.InitDirectoryConfig(*directory)

	err := envPrepare()
	if err != nil {
		logrus.Error("Env prepare error:", err.Error())
		os.Exit(1)
	}

	// start service
	var listenPort = fmt.Sprintf(":%d", *port)
	err = router.Run(listenPort)
	if err != nil {
		logrus.Error("Service start:", err.Error())
		os.Exit(2)
	}

	logrus.Info("Service end")
}

func envPrepare() error {
	// prepare
	err := service.PrepareServiceRequiredFolders()
	if err != nil {
		return err
	}

	return nil
}