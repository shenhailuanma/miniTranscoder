package main

import (
	"flag"
	"fmt"
	"github.com/shenhailuanma/miniTranscoder/router"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	port = flag.Int("p", 9000, "custom service port")
)

func main() {

	flag.Parse()

	logrus.Info("Service start")

	// start service
	var listenPort = fmt.Sprintf(":%d", *port)
	err := router.Run(listenPort)
	if err != nil {
		logrus.Error("Service start:", err.Error())
		os.Exit(1)
	}

	logrus.Info("Service end")
}
