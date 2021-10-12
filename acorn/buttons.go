package acorn

import "github.com/arskang/gomail-acorn-template/acorntypes"

type button struct {
	Params *acorntypes.ButtonParams
}

func (h HTML) NewButton(params *acorntypes.ButtonParams, btnType *string) string {
	button := &button{Params: params}
	return button.getButton(btnType)
}

func (b button) getButtonOutlined() string {
	return `
	<table cellpadding="0" cellspacing="0" role="presentation">
		<tr>
			<th style="border: 2px solid ` + b.Params.HexColor + `; border-radius: 3px; mso-padding-alt: 6px 42px 12px;">
				<a href="` + b.Params.Link + `" style="color: ` + b.Params.HexColor + `; display: inline-block; font-size: 13px; line-height: 100%; padding: 12px 42px; text-decoration: none;">` + b.Params.Text + `</a>
			</th>
		</tr>
	</table>
	`
}

func (b button) getButtonFilled() string {
	return `
	<table cellpadding="0" cellspacing="0" role="presentation">
		<tr>
			<th bgcolor="` + b.Params.HexColor + `" style="border-radius: 3px; mso-padding-alt: 6px 42px 12px;">
			<a href="` + b.Params.Link + `" style="color: #FFFFFF; display: inline-block; font-size: 13px; line-height: 100%; padding: 12px 42px; text-decoration: none;">` + b.Params.Text + `</a>
			</th>
		</tr>
	</table>
	`
}

func (b button) getButtonPill() string {
	return `
	<table cellpadding="0" cellspacing="0" role="presentation">
		<tr>
			<th bgcolor="` + b.Params.HexColor + `" style="border-radius: 50px; mso-padding-alt: 6px 42px 12px;">
			<a href="` + b.Params.Link + `" style="color: #FFFFFF; display: inline-block; font-size: 13px; line-height: 100%; padding: 12px 42px; text-decoration: none;">` + b.Params.Text + `</a>
			</th>
		</tr>
	</table>
	`
}

func (b button) getButton(btnType *string) string {
	if btnType != nil {
		switch *btnType {
		case "outlined":
			return b.getButtonOutlined()
		case "pill":
			return b.getButtonPill()
		}
	}
	return b.getButtonFilled()
}
