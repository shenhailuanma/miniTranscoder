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
	RelativePath string // relative path
	Description  string
	Publish      bool
	Snapshot     string // video snapshot id
	Custom       string // user custom data
}

const (
	JobStatusInit        = "init"
	JobStatusQueuing     = "queuing"
	JobStatusProgressing = "progressing"
	JobStatusDone        = "done"
	JobStatusError       = "error"
)

type JobUpdateRequest struct {
	Output       *string
	OutputFormat *string
	Status       *string
	SourceSize   *int64
	Progress     *int
	Command      *string
	OutputSize   *int64
	RelativePath *string // relative path
	Description  *string
	Publish      *bool
	Snapshot     *string // video snapshot id
	Custom       *string // user custom data
}
