package acornmail

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/arskang/gomail-acorn-template/acorn"
	"github.com/arskang/gomail-acorn-template/acorntypes"
)

func NewAcornEmailComponents() *acorn.HTML {
	return &acorn.HTML{}
}

func MergeVariables(temp string, variables acorntypes.AcornVariables) (string, error) {
	t, err := template.New("acorn-template").Parse(temp)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer

	err = t.Execute(&tpl, variables)
	if err != nil {
		return "", err
	}

	fmt.Println(tpl.String())

	return tpl.String(), nil
}
