package service

import "github.com/shenhailuanma/miniTranscoder/models"

func GetPlaylist() ([]models.Job, error) {
	var output = []models.Job{}

	jobs, err := GetAllJobsInfo()
	if err != nil {
		return output, err
	}

	for _, jobOne := range jobs {
		if jobOne.Publish {
			output = append(output, jobOne)
		}
	}

	return output, nil
}