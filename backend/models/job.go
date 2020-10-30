package models

type JobTable struct {
	ID         int
	Input      string
	Output     string
	Status     string
	MD5        string
	SourceName string
}

type JobStruct struct {
}
