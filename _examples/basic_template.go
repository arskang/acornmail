package examples

import (
	acornmail "github.com/arskang/gomail-acorn-template"
	"github.com/arskang/gomail-acorn-template/acornstyles"
	"github.com/arskang/gomail-acorn-template/acorntypes"
)

type Colors struct {
	Title            *acorntypes.Color
	Text             *acorntypes.Color
	BackgroundFooter *acorntypes.Color
}

func GetColors() *Colors {
	title, _ := acornstyles.NewAcornColor("#4a7cf3")
	text, _ := acornstyles.NewAcornColor("#555555")
	backgroundFooter, _ := acornstyles.NewAcornColor("#e7eef7")
	return &Colors{
		Title:            title,
		Text:             text,
		BackgroundFooter: backgroundFooter,
	}
}

func GetTitle() *acorntypes.ColumnParams {
	return &acorntypes.ColumnParams{
		Content: "<h1>{{.Title}}</h1>",
		Styles: &acorntypes.ColumnStyles{
			Align:     acornstyles.GetAligns().Center,
			TextColor: GetColors().Title,
		},
	}
}

func GetHeader() [][]*acorntypes.ColumnParams {
	acorn := acornmail.NewAcornEmailComponents()
	headerTitle := GetTitle()
	return [][]*acorntypes.ColumnParams{
		{{Content: acorn.NewSpacer()}},
		{headerTitle},
		{{Content: acorn.NewDivider(GetColors().Title)}},
	}
}

func GetButton() string {
	acorn := acornmail.NewAcornEmailComponents()
	return acorn.NewButton(&acorntypes.ButtonParams{
		Text: "{{.TxtButton}}",
		Link: "{{.Url}}",
		Styles: &acorntypes.ButtonStyles{
			FullWidth: true,
			Color:     GetColors().Title,
			Type:      acornstyles.GetTypes().Pill,
			Align:     acornstyles.GetAligns().Center,
		},
	})
}

func GetFooter() string {
	acorn := acornmail.NewAcornEmailComponents()
	row := acorn.NewRow([]*acorntypes.ColumnParams{
		{
			Content: "<b>{{.Footer}}</b>",
			Styles: &acorntypes.ColumnStyles{
				Width: acornstyles.GetWidthColumns().ThreeQuarters,
				Align: acornstyles.GetAligns().Center,
			},
		},
	})

	colors := GetColors()
	return acorn.NewAlert(&acorntypes.AlertParams{
		Content: row,
		Styles: &acorntypes.AlertStyles{
			Color:     colors.BackgroundFooter,
			TextColor: colors.Text,
		},
	})
}

func GetBasicTemplate(content string, variables acorntypes.AcornVariables, button bool) (string, error) {

	acorn := acornmail.NewAcornEmailComponents()

	rowSpacer := []*acorntypes.ColumnParams{
		{Content: acorn.NewSpacer()},
	}

	columns := append(
		GetHeader(),
		[]*acorntypes.ColumnParams{{Content: content}},
		rowSpacer,
	)

	if button {
		columns = append(
			columns,
			[]*acorntypes.ColumnParams{{Content: GetButton()}},
			rowSpacer,
		)
	}

	boilerplate := acorn.GetBoilerplate(acorntypes.AcornComponents{
		acorn.NewGrid(columns),
		GetFooter(),
	}, acornstyles.WithoutMargins())

	html, err := acornmail.MergeVariables(boilerplate, variables)
	if err != nil {
		return "", err
	}

	return html, nil
}
