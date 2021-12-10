package acorn

import (
	"github.com/arskang/acornmail/acornstyles"
	"github.com/arskang/acornmail/acorntypes"
)

type coupon struct {
	Params *acorntypes.CouponParams
}

// NewCoupon generate a new coupon html element
func (h HTML) NewCoupon(params *acorntypes.CouponParams) string {
	p := &coupon{Params: params}
	var button string
	if p.Params.Button != nil {
		button = h.NewButton(p.Params.Button)
	}
	return p.getCoupon(button)
}

// NewPromo generate a new promo html element
func (h HTML) NewPromo(items *acorntypes.PromoItems) string {
	sizes := acornstyles.GetSizes()
	colorBlack := acornstyles.GetColors().Black
	var promo, symbol, description string
	if items.Promo != nil {
		color := colorBlack
		size := sizes.Px96
		if items.Promo.Color != nil {
			color = items.Promo.Color
		}
		if items.Promo.Size != nil {
			size = items.Promo.Size
		}
		promo = `
		<th style="color: ` + color.String() + `; font-size: ` + size.String() + `; line-height: 100%; word-break: break-all;">
		` + items.Promo.Value + `
		</th>
		`
	}
	if items.Symbol != nil {
		color := colorBlack
		size := sizes.Px48
		if items.Symbol.Color != nil {
			color = items.Symbol.Color
		}
		if items.Symbol.Size != nil {
			size = items.Symbol.Size
		}
		symbol = `
		<div style="color: ` + color.String() + `; font-size: ` + size.String() + `; line-height: ` + size.String() + `;">
			` + items.Symbol.Value + `
		</div>
		`
	}
	if items.Description != nil {
		color := colorBlack
		size := sizes.Px36
		if items.Description.Color != nil {
			color = items.Description.Color
		}
		if items.Description.Size != nil {
			size = items.Description.Size
		}
		description = `
		<div style="color: ` + color.String() + `; font-size: ` + size.String() + `; line-height: ` + size.String() + `; mso-line-height-rule: exactly; mso-text-raise: 2px;">
		` + items.Description.Value + `
		</div>
		`
	}
	return `
	<table align="center" cellpadding="0" cellspacing="0" role="presentation" style="margin: 0 auto;">
		<tr>
			` + promo + `
			<th style="vertical-align: middle;">
				` + symbol + description + `
			</th>
		</tr>
	</table>
	`
}

func (c coupon) getCouponDashed(button string) string {
	acornColors := acornstyles.GetColors()
	color := acornColors.White
	borderColor := acornColors.Grey.M300
	if c.Params.Styles != nil && c.Params.Styles.Color != nil {
		color = c.Params.Styles.Color
	}
	if c.Params.Styles != nil && c.Params.Styles.BorderColor != nil {
		borderColor = c.Params.Styles.BorderColor
	}
	return `
	<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
		<tr>
			<td style="padding: 0 32px;" bgcolor="` + color.String() + `">
				<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
					<tr>
						<td class="col" align="center" width="100%">
							<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
								<tr>
									<td class="spacer py-sm-16" height="32"></td>
								</tr>
								<tr>
									<td class="px-sm-8" align="center" width="100%" style="padding: 32px; border: 4px dashed ` + borderColor.String() + `;">
										` + c.Params.Content + `
										<table cellpadding="0" cellspacing="0" role="presentation">
											<tr>
												` + button + `
											</tr>
										</table>
										<div class="spacer py-sm-8" style="line-height: 16px;">&zwnj;</div>
									</td>
								</tr>
								<tr>
									<td class="spacer py-sm-16" height="32"></td>
								</tr>
							</table>
						</td>
					</tr>
				</table>
			</td>
		</tr>
	</table>
	`
}

func (c coupon) getCouponFilled(button string) string {
	acornColors := acornstyles.GetColors()
	color := acornColors.Blue.M500
	borderColor := acornColors.White
	if c.Params.Styles != nil && c.Params.Styles.Color != nil {
		color = c.Params.Styles.Color
	}
	if c.Params.Styles != nil && c.Params.Styles.BorderColor != nil {
		borderColor = c.Params.Styles.BorderColor
	}
	return `
	<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
		<tr>
			<td class="p-sm-16" bgcolor="` + color.String() + `" width="100%" style="padding: 32px;">
				<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
					<tr>
						<td class="col" align="center" width="100%" style="border: 1px solid ` + borderColor.String() + `;">
							<table>
								<tr>
									<td class="p-sm-16" style="padding: 32px;">
										` + c.Params.Content + `
										<table cellpadding="0" cellspacing="0" role="presentation">
											<tr>
												` + button + `
											</tr>
										</table>
									</td>
								</tr>
							</table>
						</td>
					</tr>
				</table>
			</td>
		</tr>
	</table>
	`
}

func (c coupon) getCoupon(button string) string {
	if c.Params.Styles != nil && c.Params.Styles.Dashed {
		return c.getCouponDashed(button)
	}
	return c.getCouponFilled(button)
}
