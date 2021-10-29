package acorn

import (
	"github.com/arskang/gomail-acorn-template/acornstyles"
	"github.com/arskang/gomail-acorn-template/acorntypes"
)

// Generate a new content html element
func (h HTML) NewContent(params *acorntypes.ContentParams) string {
	colors := acornstyles.GetColors()
	color := colors.Grey.M800
	if params.Color != nil {
		color = params.Color
	}
	return `
	<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
		<tr>
			<td
				class="col p-sm-16"
				align="center"
				bgcolor="` + color.String() + `"
				width="100%"
				background="` + params.Image + `"
				style="background-image: url('` + params.Image + `'); background-repeat: no-repeat; background-position: center; background-size: cover;"
			>
				<div class="spacer" style="line-height: 162px; height: 162px;"></div>
					` + params.Content + `
				<div class="spacer" style="line-height: 162px; height: 162px;"></div>
			</td>
		</tr>
	</table>
	`
}
