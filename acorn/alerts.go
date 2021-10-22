package acorn

import (
	"github.com/arskang/gomail-acorn-template/acornstyles"
	"github.com/arskang/gomail-acorn-template/acorntypes"
)

type alert struct {
	Params *acorntypes.AlertParams
}

func (h HTML) NewAlert(params *acorntypes.AlertParams) string {
	a := &alert{Params: params}
	return a.getAlert()
}

func (a alert) getAlertBG(color, txtColor acorntypes.Color) string {
	return `
	<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
		<tr>
			<td bgcolor="` + color.String() + `;" style="color: ` + txtColor.String() + `; padding: 16px 32px;">
				` + a.Params.Content + `
			</td>
		</tr>
	</table>  
	`
}

func (a alert) getAlertOutlined(color, txtColor acorntypes.Color) string {
	return `
	<table cellpadding="0" cellspacing="0" role="presentation" width="100%" style="border: 2px solid ` + color.String() + `;">
		<tr>
			<td style="color: ` + txtColor.String() + `; padding: 16px 32px;">
			` + a.Params.Content + `
			</td>
		</tr>
	</table>
	`
}

func (a alert) getAlert() string {
	if a.Params != nil {
		acornColors := acornstyles.GetColors()
		color := acornColors.Blue.M800
		txtColor := acornColors.White
		if a.Params.Styles != nil {
			if a.Params.Styles.Color != nil && acornstyles.IsHexColor(string(*a.Params.Styles.Color)) {
				color = a.Params.Styles.Color
			}
			if a.Params.Styles.TextColor != nil && acornstyles.IsHexColor(string(*a.Params.Styles.TextColor)) {
				txtColor = a.Params.Styles.TextColor
			} else if a.Params.Styles.Outlined {
				txtColor = color
			}
		}

		if a.Params.Styles != nil && a.Params.Styles.Outlined {
			return a.getAlertOutlined(*color, *txtColor)
		}
		return a.getAlertBG(*color, *txtColor)
	}
	return ""
}
