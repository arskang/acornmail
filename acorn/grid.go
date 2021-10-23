package acorn

import (
	"fmt"
	"strings"

	"github.com/arskang/gomail-acorn-template/acornstyles"
	"github.com/arskang/gomail-acorn-template/acorntypes"
)

type row struct {
	Params []*acorntypes.ColumnParams
}

func (h HTML) NewRow(params []*acorntypes.ColumnParams) string {
	r := &row{Params: params}
	return r.getRow()
}

func (g row) getCol(c *acorntypes.ColumnParams) string {
	widthColumns := acornstyles.GetWidthColumns()
	aligns := acornstyles.GetAligns()
	colors := acornstyles.GetColors()
	width := widthColumns.Full
	align := aligns.Left
	txtColor := colors.Black

	var color string

	if c.Styles != nil {
		if c.Styles.Width != nil {
			width = c.Styles.Width
		}
		if c.Styles.Align != nil {
			align = c.Styles.Align
		}
		if c.Styles.Color != nil {
			color = `bgcolor="` + c.Styles.Color.String() + `"`
		}
		if c.Styles.TextColor != nil {
			txtColor = c.Styles.TextColor
		}
	}

	return fmt.Sprintf(
		`<td class="col" align="%s" width="%s" `+color+` style="color: `+txtColor.String()+`; padding: 8px; margin-left: 4px; margin-right: 4px;">%s</td>`,
		align.String(), width.String(), c.Content,
	)
}

func (g row) getRow() string {
	var columns []string
	if len(g.Params) > 0 {
		for _, column := range g.Params {
			col := g.getCol(column)
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
