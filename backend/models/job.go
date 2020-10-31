package models

// jobs table
type Job struct {
	ID         int
	Input      string
	Output     string
	Status     string
	MD5        string
	SourceName string
	SourceSize int64
	Progress   int
	Command    string
}

type JobStruct struct {
}
