package main

import (
	"flag"
	"fmt"
	"github.com/shenhailuanma/miniTranscoder/router"
	"github.com/shenhailuanma/miniTranscoder/service"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	port = flag.Int("p", 9000, "custom service port")
)

func main() {

	flag.Parse()

	logrus.Info("Service start")
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

	err = service.PrepareDatabase()
	if err != nil {
		return err
	}

	return nil
}