package runtime

import "github.com/sirupsen/logrus"

var gJobsChan chan int

func init()  {
	logrus.Info("runtime.queue init")
	gJobsChan = make(chan int, 4096)

	// todo: load not completed jobs when service start

	// push jobs to queue

	logrus.Info("runtime.queue init done")
}

func JobPush(jobID int)  {
	gJobsChan<- jobID
}

func JobPop() int {
	return <-gJobsChan
}