package acorn

import (
	"fmt"
	"strings"

	"github.com/arskang/gomail-acorn-template/acornstyles"
	"github.com/arskang/gomail-acorn-template/acorntypes"
)

type row struct {
	Params []*acorntypes.ColParams
}

func (h HTML) NewRow(params []*acorntypes.ColParams) string {
	row := &row{Params: params}
	return row.getRow()
}

func (g row) getCol(body string, styles *acorntypes.ColumnStyles) string {
	widthColumns := acornstyles.GetWidthColumns()
	width := widthColumns.Full
	if styles.Width != nil {
		width = styles.Width
	}
	aligns := acornstyles.GetAligns()
	align := aligns.Left
	if styles.Align != nil {
		align = styles.Align
	}
	return fmt.Sprintf(
		`<td class="col" align="%s" width="%s">%s</td>`,
		align.String(), width.String(), body,
	)
}

func (g row) getRow() string {
	var columns []string
	if len(g.Params) > 0 {
		for _, column := range g.Params {
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
