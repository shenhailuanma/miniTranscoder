package models

// jobs table
type Job struct {
	ID           string
	Input        string
	Output       string
	OutputFormat string
	Status       string
	MD5          string
	SourceName   string
	SourceSize   int64
	Progress     int
	Command      string
	OutputSize   int64
}

