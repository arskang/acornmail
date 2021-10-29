package acorn

import (
	"github.com/arskang/gomail-acorn-template/acornstyles"
	"github.com/arskang/gomail-acorn-template/acorntypes"
)

// Generate a new spacer html element
func (h HTML) NewSpacer() string {
	return `<div class="spacer py-sm-16" style="line-height: 32px;">&zwnj;</div>`
}

// Generate a new divider html element
func (h HTML) NewDivider(color *acorntypes.Color) string {
	acornColors := acornstyles.GetColors()
	c := acornColors.Grey.M400
	if color != nil {
		c = color
	}
	return `
	<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
		<tr>
			<td class="divider py-sm-24" style="padding: 16px;">
				<div style="background: ` + c.String() + `; height: 2px; line-height: 2px;" />
			</td>
		</tr>
	</table>
	`
}
