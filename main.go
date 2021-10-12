package acornmail

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/arskang/gomail-acorn-template/acorn"
)

func NewAcornEmailComponents() *acorn.HTML {
	return &acorn.HTML{}
}

func GetHTMLString(temp, name string, params map[string]interface{}) (string, error) {
	t, err := template.New(name).Parse(temp)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer

	err = t.Execute(&tpl, params)
	if err != nil {
		return "", err
	}

	fmt.Println(tpl.String())

	return tpl.String(), nil
}
