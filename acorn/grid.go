package acorn

import "strings"

func (h HTML) GetCol(body string, width *string) string {
	var ancho string
	switch ancho {
	case "1/4":
		ancho = "138"
	case "1/2":
		ancho = "276"
	case "3/4":
		ancho = "414"
	case "1/3":
		ancho = "184"
	case "2/3":
		ancho = "368"
	default:
		ancho = "100%"
	}
	return `<td class="col" width="` + ancho + `">` + body + `</td>`
}

func (h HTML) GetRow(columns ...string) string {
	return `
	<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
		<tr>
			<td style="padding: 0 24px;">
				<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
					<tr>
						` + strings.Join(columns, " ") + `
					</tr>
				</table>
			</td>
		</tr>
	</table>
	`
}
