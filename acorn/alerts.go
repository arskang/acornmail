package acorn

func (h HTML) GetAlertBG(content, hex string) string {
	return `
	<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
		<tr>
			<td bgcolor="` + hex + `;" style="color: #FFFFFF; padding: 16px 32px;">
				` + content + `
			</td>
		</tr>
	</table>  
	`
}

func (h HTML) GetAlertOulined(content, hex string) string {
	return `
	<table cellpadding="0" cellspacing="0" role="presentation" width="100%" style="border: 2px solid ` + hex + `;">
		<tr>
			<td style="color: ` + hex + `; padding: 16px 32px;">
			` + content + `
			</td>
		</tr>
	</table>
	`
}
