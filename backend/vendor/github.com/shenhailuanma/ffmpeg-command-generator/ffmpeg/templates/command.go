package templates

import (
	"bytes"
	"regexp"
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

	// delete '\n' '\t'
	cmdString = strings.Replace(cmdString, "\n", " ", -1)
	cmdString = strings.Replace(cmdString, "\t", " ", -1)
	cmdString = strings.TrimSpace(cmdString)

	cmdString = deleteExtraSpaces(cmdString)

	return cmdString, nil
}

func deleteExtraSpaces(input string) string {

	regstr := "\\s{2,}"
	reg := regexp.MustCompile(regstr)

	s2 := make([]byte, len(input))
	copy(s2, input)
	spc_index := reg.FindStringIndex(string(s2))
	for len(spc_index) > 0 {
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...)
		spc_index = reg.FindStringIndex(string(s2))
	}
	return string(s2)
}
