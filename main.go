package acornmail

import (
	"bytes"
	"html/template"

	"github.com/arskang/acornmail/acorn"
	"github.com/arskang/acornmail/acorntypes"
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

	return tpl.String(), nil
}
