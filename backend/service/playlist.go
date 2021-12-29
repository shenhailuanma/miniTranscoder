package service

import (
	"github.com/shenhailuanma/miniTranscoder/cache"
	"github.com/shenhailuanma/miniTranscoder/models"
	"github.com/sirupsen/logrus"
)

func GetPlaylist() ([]models.Job, error) {
	var output = []models.Job{}
	var err error

	// get from cache
	jobs, exist := cache.GetCacheJobs()
	if !exist {
		logrus.Info("GetPlaylist, not get from cache, will get data directly")
		jobs, err = GetAllJobsInfo()
		if err != nil {
			return output, err
		}

		// set cache
		cache.SetCacheJobs(jobs)
	}

	for _, jobOne := range jobs {
		if jobOne.Publish {
			output = append(output, jobOne)
		}
	}

	return output, nil
}