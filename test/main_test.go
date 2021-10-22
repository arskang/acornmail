package test

import (
	"fmt"
	"testing"

	acornmail "github.com/arskang/gomail-acorn-template"
	"github.com/arskang/gomail-acorn-template/acornstyles"
	"github.com/arskang/gomail-acorn-template/acorntypes"
)

func TestMergeVariables(t *testing.T) {

	// MergeVariables
	html, err := acornmail.MergeVariables(
		"<div>{{.Title}}</div>",
		acorntypes.AcornVariables{
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

	// Utilidades
	wColumns := acornstyles.GetWidthColumns()
	colors := acornstyles.GetColors()
	types := acornstyles.GetTypes()
	aligns := acornstyles.GetAligns()

	// NewRow
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

	alert := acorn.NewAlert(&acorntypes.AlertParams{
		Content: "Esta es una alerta",
		Styles: &acorntypes.AlertStyles{
			Outlined:  true,
			TextColor: colors.Pink.M300,
		},
	})

	buttonFilled := acorn.NewButton(&acorntypes.ButtonParams{
		Text: "Filled button",
		Link: "http://docs.thememountain.com/acorn/",
		Styles: &acorntypes.ButtonStyles{
			FullWidth: true,
		},
	})

	buttonOutlined := acorn.NewButton(&acorntypes.ButtonParams{
		Text: "Outlined button",
		Link: "http://docs.thememountain.com/acorn/",
		Styles: &acorntypes.ButtonStyles{
			Type: types.Outlined,
		},
	})

	buttonPhill := acorn.NewButton(&acorntypes.ButtonParams{
		Text: "Pill button",
		Link: "http://docs.thememountain.com/acorn/",
		Styles: &acorntypes.ButtonStyles{
			Type:  types.Pill,
			Align: aligns.Center,
		},
	})

	fmt.Println(buttonPhill)

	boilertemplate := acorn.GetBoilerPlate(
		row,
		alert,
		buttonFilled,
		buttonOutlined,
	)

	// fmt.Println(boilertemplate)
	t.Log(boilertemplate)
}
