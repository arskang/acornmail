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
	widthColumns := acornstyles.GetWidthColumns()
	colors := acornstyles.GetColors()
	types := acornstyles.GetTypes()
	aligns := acornstyles.GetAligns()

	// NewRow
	row := acorn.NewRow([]*acorntypes.ColumnParams{
		{
			Content: "1/4 de columna",
			Styles: &acorntypes.ColumnStyles{
				Width: widthColumns.Quarter,
			},
		},
		{
			Content: "1/2 de columna",
			Styles: &acorntypes.ColumnStyles{
				Width:     widthColumns.Medium,
				Color:     colors.Purple.M700,
				TextColor: colors.White,
			},
		},
		{
			Content: "1/4 de columna",
			Styles: &acorntypes.ColumnStyles{
				Width: widthColumns.Quarter,
			},
		},
	})

	// NewAlert
	alert := acorn.NewAlert(&acorntypes.AlertParams{
		Content: "Alert",
	})

	alertOutlined := acorn.NewAlert(&acorntypes.AlertParams{
		Content: "Alert",
		Styles: &acorntypes.AlertStyles{
			Outlined: true,
		},
	})

	// NewButton
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

	// NewAccordion
	accordion := acorn.NewAccordion([]*acorntypes.AccordionParams{
		{
			Title:   "Panel 1",
			Content: "Lorem ipsum dolor sit amet, consectetur adipisicing elit.",
		},
		{
			Title:   "Panel 2",
			Content: "Lorem ipsum dolor sit amet, consectetur adipisicing elit.",
			Styles: &acorntypes.AccordionStyles{
				Color:        colors.Cyan.M300,
				TitleColor:   colors.White,
				ContentColor: colors.Cyan.M300,
			},
		},
	})

	// NewSpacer
	spacer := acorn.NewSpacer()

	// NewDivider
	divider := acorn.NewDivider(colors.DeepPurple.M700)

	// Label
	labelFilled := acorn.NewLabel(&acorntypes.LabelParams{
		Text: "filled label",
	})

	labelOutlined := acorn.NewLabel(&acorntypes.LabelParams{
		Text: "outlined label",
		Styles: &acorntypes.LabelStyles{
			Type: types.Outlined,
		},
	})

	// Content
	content := acorn.NewContent(&acorntypes.ContentParams{
		Content: row,
		Image:   "",
	})

	// withoutMargins := true
	// b := acorn.GetBoilerplate(acorntypes.AcornComponents{content}, &withoutMargins)

	fmt.Println(accordion)

	boilertemplate := acorn.GetBoilerplate(acorntypes.AcornComponents{
		row,
		alert,
		buttonFilled,
		buttonOutlined,
		buttonPhill,
		accordion,
		spacer,
		divider,
		alertOutlined,
		labelFilled,
		labelOutlined,
		content,
	}, nil)

	t.Log(boilertemplate)
}

func TestBasicExample(t *testing.T) {

	// acorn := acornmail.NewAcornEmailComponents()

	// widthColumns := acornstyles.GetWidthColumns()
	// colors := acornstyles.GetColors()
	// types := acornstyles.GetTypes()
	// aligns := acornstyles.GetAligns()

}
