package ffmpeg

import "github.com/shenhailuanma/ffmpeg-command-generator/ffmpeg/templates"

func FFmpegSnapshot(request FFmpegTranscodeRequest) (string, error) {
	return templates.GenerateCommand("snapshot", templates.SnapshotTemplate, request)
}