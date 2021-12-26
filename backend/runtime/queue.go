package runtime

import "github.com/sirupsen/logrus"

var gJobsChan chan string

func init()  {
	logrus.Info("runtime.queue init")
	gJobsChan = make(chan string, 4096)

	// todo: load not completed jobs when service start

	// push jobs to queue

	logrus.Info("runtime.queue init done")
}

func JobPush(jobID string)  {
	gJobsChan<- jobID
}

func JobPop() string {
	return <-gJobsChan
}