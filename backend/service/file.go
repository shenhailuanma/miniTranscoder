package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shenhailuanma/miniTranscoder/config"
	"github.com/shenhailuanma/miniTranscoder/models"
	"github.com/shenhailuanma/miniTranscoder/utils"
	"github.com/sirupsen/logrus"
	"sort"
	"strconv"
)

/**
File tree definition:

/ -- ${BaseFolder}
	 	/-- upload
			-- ${SourceVideo1}
			-- ${SourceVideo1}
				...
			-- ${SourceVideoN}
        /-- data
			/-- ${JobNumber}
  					-- ${SourceFile}.${Format}
					-- job.json
					-- ${CustomName}.mp4
					-- ${CustomName}.m3u8
					/-- m3u8
					/-- images
			/-- 2
			  ...
			/-- 102
  				-- source.mp4
				-- job.json
				-- ${CustomName}.mp4
				-- ${CustomName}.m3u8
				/-- m3u8

*/

func InitJob(job models.Job) (string, error) {

	// get all job folders and get the latest number
	jobFolders, err := GetJobFolders()
	if err != nil {
		return "", err
	}

	var jobID = "1" // default

	// the latest number
	if len(jobFolders) > 0 {
		numberString := jobFolders[len(jobFolders)-1]
		numberValue, err := strconv.Atoi(numberString)
		if err != nil {
			return "", errors.New("[InitJob]Parse jobID:" + numberString + ", error:" + err.Error())
		}
		jobID = fmt.Sprintf("%d", numberValue+1)
	}

	// create job folder
	jobFolder := fmt.Sprintf("%s/%s", config.ConfigDataOutputPath, jobID)
	err = utils.CreatePath(jobFolder)
	if err != nil {
		return "", errors.New("[InitJob]CreatePath error:" + err.Error())
	}

	// create job config file
	err = SetJobConfig(jobID, job)
	if err != nil {
		return "", errors.New("[InitJob]Create Job config file error:" + err.Error())
	}

	return jobID, nil
}

func GetAllJobsInfo() ([]models.Job, error) {
	var output = []models.Job{}

	// list all job date folders
	jobFolders, err := GetJobFolders()
	if err != nil {
		return output, err
	}

	// get jobs info
	for _, jobFolderOne := range jobFolders {
		jobOne, err := GetJobConfig(jobFolderOne)
		if err != nil {
			logrus.Warn("GetAllJobsInfo, GetJobConfig:", jobFolderOne, ", error:", err.Error())
			continue
		}

		jobOne.ID = jobFolderOne
		output = append(output, jobOne)
	}

	return output, err
}

func GetJobConfig(jobID string) (models.Job, error) {
	var output = models.Job{}

	jobConfigFile := fmt.Sprintf("%s/%s/job.json", config.ConfigDataOutputPath, jobID)

	// check exist
	exist, _ := utils.PathExists(jobConfigFile)
	if !exist {
		return output, errors.New("Job not exist")
	}

	// read file
	fileData, err := utils.ReadFile(jobConfigFile)
	if err != nil {
		return output, errors.New("Read jobConfig:" + jobConfigFile + " error:" + err.Error())
	}

	err = json.Unmarshal(fileData, &output)
	if err != nil {
		return output, errors.New("Unmarshal jobConfig error:" + err.Error())
	}

	return output, nil
}

func SetJobConfig(jobID string, job models.Job) error {
	jobConfigFile := fmt.Sprintf("%s/%s/job.json", config.ConfigDataOutputPath, jobID)

	job.ID = jobID

	// create job file
	jobData, err := json.MarshalIndent(job, "", "  ")
	if err != nil {
		return errors.New("JobData Marshal error:" + err.Error())
	}

	return utils.WriteBytesToFile(jobConfigFile, jobData)
}

func UpdateJobStatus(jobID string, status string) error {
	jobInfo, err := GetJobConfig(jobID)
	if err != nil {
		return err
	}

	jobInfo.Status = status

	return SetJobConfig(jobID, jobInfo)
}

func UpdateJobProgress(jobID string, progress int) error {
	jobInfo, err := GetJobConfig(jobID)
	if err != nil {
		return err
	}

	jobInfo.Progress = progress

	return SetJobConfig(jobID, jobInfo)
}

func UpdateJobInfo(jobID string, request models.JobUpdateRequest) error {
	jobInfo, err := GetJobConfig(jobID)
	if err != nil {
		return err
	}

	if request.Output != nil {
		jobInfo.Output = *request.Output
	}
	if request.OutputFormat != nil {
		jobInfo.OutputFormat = *request.OutputFormat
	}
	if request.Status != nil {
		jobInfo.Status = *request.Status
	}
	if request.SourceSize != nil {
		jobInfo.SourceSize = *request.SourceSize
	}
	if request.Progress != nil {
		jobInfo.Progress = *request.Progress
	}

	if request.Command != nil {
		jobInfo.Command = *request.Command
	}
	if request.OutputSize != nil {
		jobInfo.OutputSize = *request.OutputSize
	}
	if request.RelativePath != nil {
		jobInfo.RelativePath = *request.RelativePath
	}

	if request.Description != nil {
		jobInfo.Description = *request.Description
	}
	if request.Publish != nil {
		jobInfo.Publish = *request.Publish
	}
	if request.Snapshot != nil {
		jobInfo.Snapshot = *request.Snapshot
	}
	if request.Custom != nil {
		jobInfo.Custom = *request.Custom
	}

	return SetJobConfig(jobID, jobInfo)
}

func GetJobFolders() ([]string, error) {
	var output = []string{}
	// list all job folders
	jobFolders, err := utils.GetDirFilenames(config.ConfigDataOutputPath)
	if err != nil {
		return output, err
	}

	var jobIDs = []int{}
	for _, jobOne := range jobFolders {
		jobIDOne, err := strconv.Atoi(jobOne)
		if err == nil && jobIDOne > 0 {
			jobIDs = append(jobIDs, jobIDOne)
		}
	}

	// sort
	sort.Ints(jobIDs)
	for _, jobIDOne := range jobIDs {
		output = append(output, fmt.Sprintf("%d", jobIDOne))
	}

	return output, nil
}

func JobOutputPath(jobID string, format string) string {
	return fmt.Sprintf("%s/%s/video.%s", config.ConfigDataOutputPath, jobID, format)
}

func JobRelativePath(jobID string, format string) string {
	return fmt.Sprintf("%s/%s/video.%s", config.ConfigVodFolder, jobID, format)
}

func GetJobOutputFileSize(jobID string) (int64, error) {
	jobInfo, err := GetJobConfig(jobID)
	if err != nil {
		return 0, err
	}

	return utils.FileSize(jobInfo.Output), nil
}

func SyncJobOutputFileSize(jobID string) error {
	jobInfo, err := GetJobConfig(jobID)
	if err != nil {
		return err
	}

	size, err := GetJobOutputFileSize(jobID)
	if err != nil {
		return err
	}

	jobInfo.OutputSize = size

	return SetJobConfig(jobID, jobInfo)
}
