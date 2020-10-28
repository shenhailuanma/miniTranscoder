package ffmpeg

type FFmpegGlobalParams struct {
	Overwrite bool `json:"overwrite"`
	NoStream  bool `json:"noStream"` // -vn, as an input option, blocks all video streams
}
