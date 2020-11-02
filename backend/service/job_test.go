package service

import "testing"

func Test_parseFFmpegLogProgress(t *testing.T)  {
	progress := parseFFmpegLogProgress("../miniTranscoder/data/output/26.log")
	t.Log("progress:", progress)
}