package acorn

type alert struct {
	Content, HexColor string
	Outlined          bool
}

func (h HTML) NewAlert(content, hex string, outlined *bool) string {
	var out bool
	if outlined != nil && *outlined {
		out = true
	}
	alert := &alert{
		Content:  content,
		HexColor: hex,
		Outlined: out,
	}
	return alert.getAlert()
}

func (a alert) getAlertBG() string {
	return `
	<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
		<tr>
			<td bgcolor="` + a.HexColor + `;" style="color: #FFFFFF; padding: 16px 32px;">
				` + a.Content + `
			</td>
		</tr>
	</table>  
	`
}

func (a alert) getAlertOutlined() string {
	return `
	<table cellpadding="0" cellspacing="0" role="presentation" width="100%" style="border: 2px solid ` + a.HexColor + `;">
		<tr>
			<td style="color: ` + a.HexColor + `; padding: 16px 32px;">
			` + a.Content + `
			</td>
		</tr>
	</table>
	`
}

func (a alert) getAlert() string {
	if a.Outlined {
		return a.getAlertOutlined()
	}
	return a.getAlertBG()
}
