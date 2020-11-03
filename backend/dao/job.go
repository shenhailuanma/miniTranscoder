package dao

import "github.com/shenhailuanma/miniTranscoder/models"

func PrepareJobsTable() error {

	db, err := DatabaseOpen()
	if err != nil {
		return err
	}

	return db.AutoMigrate(&models.Job{})
}

func GetJobs(page int, size int) ([]models.Job, error) {
	var jobs = []models.Job{}

	db, err := DatabaseOpen()
	if err != nil {
		return jobs, err
	}

	if page < 0 {
		page = 0
	}

	if size <= 0 && size != -1 {
		size = 15
	}

	var offset = 0
	if page > 0 && size > 0{
		offset = size * page
	}

	err = db.Table("jobs").Order("id desc").Offset(offset).Limit(size).Find(&jobs).Error

	return jobs, err
}


func GetJobsCount() (int64, error) {
	var count int64 = 0

	db, err := DatabaseOpen()
	if err != nil {
		return count, err
	}

	err = db.Table("jobs").Count(&count).Error

	return count, err
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

func UpdateJobStatus(jobID int, status string) error  {
	db, err := DatabaseOpen()
	if err != nil {
		return err
	}
	return db.Table("jobs").Where("id=?", jobID).Update("status", status).Error
}