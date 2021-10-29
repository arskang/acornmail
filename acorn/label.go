package acorn

import (
	"github.com/arskang/gomail-acorn-template/acornstyles"
	"github.com/arskang/gomail-acorn-template/acorntypes"
)

type label struct {
	Params *acorntypes.LabelParams
}

// Generate a new label html element
func (h HTML) NewLabel(params *acorntypes.LabelParams) string {
	l := label{Params: params}
	return l.getLabel()
}

func (l label) getLabelFilled(params *acorntypes.LabelParams, color, textColor *acorntypes.Color) string {
	return `
	<span style="border-width: 2px 4px; mso-border-width-alt: 4px; border-style: solid; border-color: ` + color.String() + `; background-color: ` + color.String() + `; border-radius: 3px; color: ` + textColor.String() + `; font-size: 75%; line-height: 100%; mso-line-height-rule: exactly;">` + l.Params.Text + `</span>
	`
}

func (l label) getLabelOutlined(params *acorntypes.LabelParams, color, textColor *acorntypes.Color) string {
	return `
	<span style="padding: 1px 4px; mso-padding-alt: 4px; border: 1px solid ` + color.String() + `; border-radius: 3px; color: ` + textColor.String() + `; font-size: 75%; line-height: 100%; mso-line-height-rule: exactly;">` + l.Params.Text + `</span>
	`
}

func (l label) getLabel() string {
	if l.Params != nil {

		acornColors := acornstyles.GetColors()
		acornTypes := acornstyles.GetTypes()

		color := acornColors.Green.M800
		txtColor := acornColors.White

		if l.Params.Styles != nil {
			if l.Params.Styles.Color != nil && acornstyles.IsHexColor(string(*l.Params.Styles.Color)) {
				color = l.Params.Styles.Color
			}
			if l.Params.Styles.TextColor != nil && acornstyles.IsHexColor(string(*l.Params.Styles.TextColor)) {
				txtColor = l.Params.Styles.TextColor
			} else if l.Params.Styles.Type != nil && *l.Params.Styles.Type == *acornTypes.Outlined {
				txtColor = color
			}
		}

		if l.Params.Styles != nil && l.Params.Styles.Type != nil {
			switch *l.Params.Styles.Type {
			case *acornTypes.Outlined:
				return l.getLabelOutlined(l.Params, color, txtColor)
			}
		}
		if l.Params.Styles != nil && l.Params.Styles.Outlined {
			return l.getLabelOutlined(l.Params, color, txtColor)
		}
		return l.getLabelFilled(l.Params, color, txtColor)
	}
	return ""
}
