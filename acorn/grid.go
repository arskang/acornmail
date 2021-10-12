package acorn

import (
	"strings"

	"github.com/arskang/gomail-acorn-template/acorntypes"
)

type row struct {
	Columns []*acorntypes.Col
}

func (h HTML) NewRow(columns []*acorntypes.Col) string {
	row := &row{Columns: columns}
	return row.getRow()
}

func (g row) getCol(body string, width *string) string {
	var ancho string
	switch *width {
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

func (g row) getRow() string {
	var columns []string
	if len(g.Columns) > 0 {
		for _, column := range g.Columns {
			col := g.getCol(column.Content, column.Width)
			columns = append(columns, col)
		}
	}
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
