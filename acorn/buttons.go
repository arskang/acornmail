package acorn

import (
	"github.com/arskang/gomail-acorn-template/acornstyles"
	"github.com/arskang/gomail-acorn-template/acorntypes"
)

type button struct {
	Params *acorntypes.ButtonParams
}

func (h HTML) NewButton(params *acorntypes.ButtonParams) string {
	b := &button{Params: params}
	return b.getButton()
}

func (b button) getButtonOutlined(color, txtColor acorntypes.Color, align acorntypes.Align, fullWidth string) string {
	return `
	<table ` + fullWidth + ` align="` + align.String() + `" cellpadding="0" cellspacing="0" role="presentation">
		<tr>
			<th style="border: 2px solid ` + color.String() + `; border-radius: 3px; mso-padding-alt: 6px 42px 12px;">
				<a href="` + b.Params.Link + `" style="color: ` + txtColor.String() + `; display: inline-block; font-size: 13px; line-height: 100%; padding: 12px 42px; text-decoration: none;">` + b.Params.Text + `</a>
			</th>
		</tr>
	</table>
	`
}

func (b button) getButtonFilled(color, txtColor acorntypes.Color, align acorntypes.Align, fullWidth string) string {
	return `
	<table ` + fullWidth + ` align="` + align.String() + `" cellpadding="0" cellspacing="0" role="presentation">
		<tr>
			<th bgcolor="` + color.String() + `" style="border-radius: 3px; mso-padding-alt: 6px 42px 12px;">
			<a href="` + b.Params.Link + `" style="color: ` + txtColor.String() + `; display: inline-block; font-size: 13px; line-height: 100%; padding: 12px 42px; text-decoration: none;">` + b.Params.Text + `</a>
			</th>
		</tr>
	</table>
	`
}

func (b button) getButtonPill(color, txtColor acorntypes.Color, align acorntypes.Align, fullWidth string) string {
	return `
	<table ` + fullWidth + ` align="` + align.String() + `" cellpadding="0" cellspacing="0" role="presentation">
		<tr>
			<th bgcolor="` + color.String() + `" style="border-radius: 50px; mso-padding-alt: 6px 42px 12px;">
			<a href="` + b.Params.Link + `" style="color: ` + txtColor.String() + `; display: inline-block; font-size: 13px; line-height: 100%; padding: 12px 42px; text-decoration: none;">` + b.Params.Text + `</a>
			</th>
		</tr>
	</table>
	`
}

func (b button) getButton() string {
	if b.Params != nil {

		acornColors := acornstyles.GetColors()
		acornTypes := acornstyles.GetTypes()
		aligns := acornstyles.GetAligns()

		color := acornColors.Green.M800
		txtColor := acornColors.White
		align := aligns.Left
		var fullWidth string

		if b.Params.Styles != nil {
			if b.Params.Styles.Color != nil && acornstyles.IsHexColor(string(*b.Params.Styles.Color)) {
				color = b.Params.Styles.Color
			}
			if b.Params.Styles.TextColor != nil && acornstyles.IsHexColor(string(*b.Params.Styles.TextColor)) {
				txtColor = b.Params.Styles.TextColor
			} else if b.Params.Styles.Type != nil && *b.Params.Styles.Type == *acornTypes.Outlined {
				txtColor = color
			}
			if b.Params.Styles.Align != nil {
				align = b.Params.Styles.Align
			}
			if b.Params.Styles.FullWidth {
				fullWidth = `width="100%"`
			}
		}

		if b.Params.Styles != nil && b.Params.Styles.Type != nil {
			switch *b.Params.Styles.Type {
			case *acornTypes.Outlined:
				return b.getButtonOutlined(*color, *txtColor, *align, fullWidth)
			case *acornTypes.Pill:
				return b.getButtonPill(*color, *txtColor, *align, fullWidth)
			}
		}
		return b.getButtonFilled(*color, *txtColor, *align, fullWidth)
	}
	return ""
}
