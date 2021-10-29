package acorn

import (
	"strings"

	"github.com/arskang/gomail-acorn-template/acornstyles"
	"github.com/arskang/gomail-acorn-template/acorntypes"
)

type accordion struct {
	Params []*acorntypes.AccordionParams
}

// Generate a new accordion html element
func (h HTML) NewAccordion(params []*acorntypes.AccordionParams) string {
	a := &accordion{Params: params}
	return a.getAccordion()
}

func (a accordion) getElement(p *acorntypes.AccordionParams, withoutSpacer bool) string {

	acornColors := acornstyles.GetColors()

	color := acornColors.Grey.M400
	titleColor := acornColors.Grey.M900
	contentColor := acornColors.Grey.M900

	if p.Styles != nil {
		if p.Styles.Color != nil && acornstyles.IsHexColor(string(*p.Styles.Color)) {
			color = p.Styles.Color
		}
		if p.Styles.TitleColor != nil && acornstyles.IsHexColor(string(*p.Styles.TitleColor)) {
			titleColor = p.Styles.TitleColor
		}
		if p.Styles.ContentColor != nil && acornstyles.IsHexColor(string(*p.Styles.ContentColor)) {
			contentColor = p.Styles.ContentColor
		}
	}

	spacer := `<div class="spacer py-sm-16" style="line-height: 32px;">&zwnj;</div>`
	if withoutSpacer {
		spacer = ""
	}

	return `
	<div>
		<a class="toggle-trigger" style="text-decoration: none;">
			<button class="toggle-trigger" style="background-color:` + color.String() + `; margin: 0; padding: 0; display: block; width: 100%; text-align: left; border: none; outline: none; font-size: 13px;">
				<table bgcolor="` + color.String() + `" cellpadding="0" cellspacing="0" role="presentation" width="100%">
					<tr>
						<td style="padding: 16px; color: ` + titleColor.String() + `;">` + p.Title + `</td>
					</tr>
				</table>
			</button>
		</a>
		<div class="toggle-content">
			<table cellpadding="0" cellspacing="0" width="100%" role="presentation" style="border: 1px solid ` + color.String() + `;">
				<tr>
					<td style="padding: 16px; color: ` + contentColor.String() + `;">
						` + p.Content + `
					</td>
				</tr>
			</table>
		</div>
	</div>
	` + spacer
}

func (a accordion) getAccordion() string {
	var elements []string
	if len(a.Params) > 0 {
		for i, p := range a.Params {
			element := a.getElement(p, i == len(a.Params)-1)
			elements = append(elements, element)
		}
	}
	return `
	<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
		<tr>
			<td style="padding: 0 24px;">
				<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
					<tr>
						<td class="col" width="100%" style="padding: 0 8px;">
							` + strings.Join(elements, " ") + `
						</td>
					</tr>
				</table>
			</td>
		</tr>
	</table>
	`
}
