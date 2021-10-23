package acorn

import (
	"strings"

	"github.com/arskang/gomail-acorn-template/acornstyles"
	"github.com/arskang/gomail-acorn-template/acorntypes"
)

type testimonail struct {
	Params *acorntypes.TestimonialParams
}

func (h HTML) NewTestimonial(params *acorntypes.TestimonialParams) string {
	t := testimonail{Params: params}
	return t.getTestimonial()
}

func (t testimonail) getAuthor() string {
	if t.Params != nil && strings.TrimSpace(t.Params.Author) != "" {
		colors := acornstyles.GetColors()
		return `<small style="color: ` + colors.Grey.M500.String() + `; text-transform: uppercase;">` + t.Params.Author + `</small>`
	}
	return ""
}

func (t testimonail) getTestimonialAvatar(align *acorntypes.Align) string {
	return `
	<td class="px-sm-16" style="padding: 0 100px;">
		<div class="spacer" style="line-height: 40px; height: 40px!important;">&zwnj;</div>
		<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
			<tr>
				<td class="col px-sm-16" align="` + align.String() + `" width="100%">
					<img src="` + t.Params.Styles.Image + `" width="50" alt="Avatar">
					<h3 style="font-weight: 300; font-style: italic;">` + t.Params.Testimonial + `</h3>
					` + t.getAuthor() + `
				</td>
			</tr>
		</table>
		<div class="spacer" style="line-height: 40px; height: 40px!important;">&zwnj;</div>
	</td>
	`
}

func (t testimonail) getTestimonialBorder(align *acorntypes.Align) string {

	colors := acornstyles.GetColors()
	borderColor := colors.Grey.M300

	if t.Params.Styles != nil && t.Params.Styles.BorderColor != nil {
		borderColor = t.Params.Styles.BorderColor
	}

	return `
	<td class="px-sm-16" style="padding: 0 60px;">
		<div class="spacer" style="line-height: 40px; height: 40px;">&zwnj;</div>
		<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
			<tr>
				<td class="col" width="100%">
					<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
						<tr>
							<td class="pl-sm-16" style="border-left: 2px solid ` + borderColor.String() + `; padding-left: 30px;" align="` + align.String() + `">
							<h3 style="font-weight: 500; font-style: italic; margin-top: 0;">` + t.Params.Testimonial + `</h3>
							` + t.getAuthor() + `
						</td>
						</tr>
					</table>
				</td>
			</tr>
		</table>
		<div class="spacer" style="line-height: 40px; height: 40px;">&zwnj;</div>
	</td>
	`
}

func (t testimonail) getTestimonialIcon(align *acorntypes.Align) string {
	return `
	<td style="padding: 0 24px;">
		<div class="spacer" style="line-height: 40px; height: 40px!important;">&zwnj;</div>
		<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
			<tr>
				<td class="col" width="100%">
					<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
						<tr>
							<td style="padding: 0 8px;" align="` + align.String() + `">
								<img src="` + IMAGE_TESTIMONIAL + `" alt="Quote" width="33">
								<h3 style="font-weight: 500; font-style: italic;">` + t.Params.Testimonial + `</h3>
								` + t.getAuthor() + `
							</td>
						</tr>
					</table>
				</td>
			</tr>
		</table>
		<div class="spacer" style="line-height: 40px; height: 40px!important;">&zwnj;</div>
	</td>
	`
}

func (t testimonail) getTestimonial() string {
	if t.Params != nil {

		var content string

		aligns := acornstyles.GetAligns()

		align := aligns.Left

		if t.Params.Styles != nil && t.Params.Styles.Align != nil {
			align = t.Params.Styles.Align
		}

		if t.Params.Icon {
			content = t.getTestimonialIcon(align)
		} else if t.Params.Styles != nil && strings.TrimSpace(t.Params.Styles.Image) != "" {
			content = t.getTestimonialAvatar(align)
		} else {
			content = t.getTestimonialBorder(align)
		}

		return `
		<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
			<tr>
				` + content + `
			</tr>
		</table>
		`
	}
	return ""
}
