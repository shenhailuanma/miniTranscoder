package templates


const SnapshotTemplate = `
{{- range $inputIndex, $inputOne := .Inputs}}
-i {{$inputOne}} 
{{- end }}

{{- range $outputIndex, $outputOne := .Outputs}}
	{{- range $streamIndex, $streamOne := $outputOne.Streams}}
		{{- if eq "video" $streamOne.Kind}}
			{{- if $streamOne.Video.Width}}
				{{- if $streamOne.Video.Height}}
					-s {{$streamOne.Video.Width}}x{{$streamOne.Video.Height}}
				{{- end }}
			{{- end }}
		{{- end }}
    {{- end }}
-f image2 -vframes 1
		-y {{$outputOne.Output}}
{{- end }}
`