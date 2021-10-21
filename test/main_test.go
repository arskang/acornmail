package test

import (
	"fmt"
	"testing"

	acornmail "github.com/arskang/gomail-acorn-template"
	"github.com/arskang/gomail-acorn-template/acornstyles"
	"github.com/arskang/gomail-acorn-template/acorntypes"
)

func TestGetHTMLString(t *testing.T) {

	// GetHTMLString
	html, err := acornmail.GetHTMLString(
		"<div>{{.Title}}</div>",
		"Ejemplo",
		map[string]interface{}{
			"Title": "Hola mundo",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(html)

}

func TestExample(t *testing.T) {
	// Componentes
	acorn := acornmail.NewAcornEmailComponents()

	// NewRow
	wColumns := acornstyles.GetWidthColumns()
	row := acorn.NewRow([]*acorntypes.ColParams{
		{
			Content: "1/4 de columna",
			Styles: &acorntypes.ColumnStyles{
				Width: wColumns.Quarter,
			},
		},
		{
			Content: "1/2 de columna",
			Styles: &acorntypes.ColumnStyles{
				Width: wColumns.Medium,
			},
		},
		{
			Content: "1/4 de columna",
			Styles: &acorntypes.ColumnStyles{
				Width: wColumns.Quarter,
			},
		},
	})

	boilertemplate := acorn.GetBoilerPlate(
		row,
	)

	fmt.Println(boilertemplate)
	t.Log(boilertemplate)
}
