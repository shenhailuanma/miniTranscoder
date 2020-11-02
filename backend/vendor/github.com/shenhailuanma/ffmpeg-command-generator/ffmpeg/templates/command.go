package templates

import (
	"bytes"
	"strings"
	"text/template"
)

func GenerateCommand(name string, templateText string, data interface{}) (string, error) {

	tmpl := template.New(name).Funcs(template.FuncMap{})

	tmpl, err := tmpl.Parse(templateText)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	var cmdString = buf.String()

	// split by '\n'
	cmdStringList := strings.Split(cmdString, "\n")
	var cmdStringOutput = ""
	for _, cmdStringOne := range cmdStringList {
		// trim space and connect all strings
		cmdStringOutput = cmdStringOutput + " " + strings.TrimSpace(cmdStringOne)
	}

	return cmdStringOutput, nil
}
