package acorn

import (
	"github.com/arskang/gomail-acorn-template/acornstyles"
	"github.com/arskang/gomail-acorn-template/acorntypes"
)

type coupon struct {
	Params *acorntypes.CouponParams
}

// Generate a new coupon html element
func (h HTML) NewCoupon(params *acorntypes.CouponParams) string {
	return ""
}

// Generate a new promo html element
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
	return `
	<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
		<tr>
			<td style="padding: 0 32px;">
				<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
					<tr>
						<td class="col" align="center" width="100%">
							<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
								<tr>
									<td class="spacer py-sm-16" height="32"></td>
								</tr>
								<tr>
									<td class="px-sm-8" align="center" width="100%" style="padding: 32px; border: 4px dashed #CCCCCC; color: #000000;">
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