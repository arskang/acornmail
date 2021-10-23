package acorn

import (
	"strings"

	"github.com/arskang/gomail-acorn-template/acornstyles"
	"github.com/arskang/gomail-acorn-template/acorntypes"
)

func (h HTML) NewImage(params *acorntypes.ImageParams) string {
	alt := "Image"
	widthColumn := acornstyles.GetWidthColumns()
	width := widthColumn.Full
	if params.WidthColumn != nil {
		width = params.WidthColumn
	}
	if strings.TrimSpace(params.Alt) != "" {
		alt = params.Alt
	}
	return `
	<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
		<tr>
			<td class="col" align="center" width="` + width.String() + `" style="padding: 0 8px;">
				<img src="` + params.Image + `" alt="` + alt + `" width="` + width.String() + `">
			</td>
		</tr>
	</table>
	`
}
