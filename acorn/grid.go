package acorn

import (
	"fmt"
	"strings"

	"github.com/arskang/gomail-acorn-template/acorntypes"
)

type row struct {
	Columns []*acorntypes.ColParams
}

func (h HTML) NewRow(columns []*acorntypes.ColParams) string {
	row := &row{Columns: columns}
	return row.getRow()
}

func (g row) getCol(body string, styles *acorntypes.ColumnStyles) string {
	width := "100%"
	if styles.Width != nil {
		width = string(*styles.Width)
	}
	align := "left"
	if styles.Align != nil {
		width = string(*styles.Align)
	}
	return fmt.Sprintf(
		`<td class="col" align:"%s" width="%s">%s</td>`,
		align, width, body,
	)
}

func (g row) getRow() string {
	var columns []string
	if len(g.Columns) > 0 {
		for _, column := range g.Columns {
			col := g.getCol(column.Content, column.Styles)
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
