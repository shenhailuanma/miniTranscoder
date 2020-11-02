package templates


const TranscodeTemplate = `
{{- range $inputIndex, $inputOne := .Inputs}}
-i '{{$inputOne}}'
{{- end }}

{{- range $outputIndex, $outputOne := .Outputs}}
	{{- range $streamIndex, $streamOne := $outputOne.Streams}}
		{{- if eq "video" $streamOne.Kind}}
			{{- if $streamOne.Video.Codec}}
		 		-c:v {{$streamOne.Video.Codec}}
			{{- end }}
			{{- if $streamOne.Video.Preset}}
		 		-preset {{$streamOne.Video.Preset}}
			{{- end }}
			{{- if $streamOne.Video.Width}}
				{{- if $streamOne.Video.Height}}
					-s {{$streamOne.Video.Width}}x{{$streamOne.Video.Height}}
				{{- end }}
			{{- end }}
			{{- if $streamOne.Video.Fps}}
		 		-r {{$streamOne.Video.Fps}}
			{{- end }}
			{{- if $streamOne.Video.Bitrate}}
		 		-b:v {{$streamOne.Video.Bitrate}}
			{{- end }}
		{{- end }}
		{{- if eq "audio" $streamOne.Kind}}
			{{- if $streamOne.Audio.Codec}}
		 		-c:a {{$streamOne.Audio.Codec}}
			{{- end }}
			{{- if $streamOne.Audio.Channels}}
		 		-ac {{$streamOne.Audio.Channels}}
			{{- end }}
		{{- end }}
    {{- end }}
	{{- if $outputOne.Format}}
		-f {{$outputOne.Format}}
	{{- end}}
		-y '{{$outputOne.Output}}'
{{- end }}
`