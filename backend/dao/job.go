package dao

import "github.com/shenhailuanma/miniTranscoder/models"

func PrepareJobsTable() error {

	db, err := DatabaseOpen()
	if err != nil {
		return err
	}

	return db.AutoMigrate(&models.Job{})
}

func GetJobs() ([]models.Job, error) {
	var jobs = []models.Job{}

	db, err := DatabaseOpen()
	if err != nil {
		return jobs, err
	}

	err = db.Table("jobs").Find(&jobs).Error

	return jobs, err
}

func CreateJob(job models.Job) (int, error) {
	var jobID = 0

	db, err := DatabaseOpen()
	if err != nil {
		return jobID, err
	}

	err = db.Table("jobs").Create(&job).Error
	return job.ID, err
}

func GetJobInfo(jobID int) (models.Job, error) {
	var job = models.Job{}

	db, err := DatabaseOpen()
	if err != nil {
		return job, err
	}

	err = db.Table("jobs").Where("id=?", jobID).First(&job).Error
	return job, err
}

func UpdateJobProgress(jobID, progress int) error  {
	db, err := DatabaseOpen()
	if err != nil {
		return err
	}
	return db.Table("jobs").Where("id=?", jobID).Update("progress", progress).Error
}