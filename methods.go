package acornmail

import (
	"bytes"
	"fmt"
	"html/template"
)

func (ae acornEmail) GetHTMLString(temp string) (string, error) {
	t, err := template.New("template").Parse(temp)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer

	err = t.Execute(&tpl, nil)
	if err != nil {
		return "", err
	}

	fmt.Println(tpl.String())

	return tpl.String(), nil
}

func (ae acornEmail) GetTemplate(body string) string {
	return ae.HTML.GetBoilerPlate(body)
}
