package test

import (
	"fmt"
	"testing"

	"github.com/arskang/acornmail"
	"github.com/arskang/acornmail/acornstyles"
	"github.com/arskang/acornmail/acorntypes"
)

func TestMergeVariables(t *testing.T) {

	// MergeVariables
	html, err := acornmail.MergeVariables(
		"<div>{{.Title}}</div>",
		acorntypes.AcornVariables{
			"Title": "Hola mundo",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(html)

}

func TestExample(t *testing.T) {
	// Components
	acorn := acornmail.NewAcornEmailComponents()

	fmt.Println(acornstyles.RGX_HEXCOLOR)
	fmt.Println(acornstyles.RGX_URL)

	// Utilidades
	widthColumns := acornstyles.GetWidthColumns()
	colors := acornstyles.GetColors()
	types := acornstyles.GetTypes()
	aligns := acornstyles.GetAligns()

	// NewRow
	row := acorn.NewRow([]*acorntypes.ColumnParams{
		{
			Content: "1/4 de columna",
			Styles: &acorntypes.ColumnStyles{
				Width: widthColumns.Quarter,
			},
		},
		{
			Content: "1/2 de columna",
			Styles: &acorntypes.ColumnStyles{
				Width:     widthColumns.Medium,
				Color:     colors.Purple.M700,
				TextColor: colors.White,
			},
		},
		{
			Content: "1/4 de columna",
			Styles: &acorntypes.ColumnStyles{
				Width: widthColumns.Quarter,
			},
		},
	})

	// NewAlert
	alert := acorn.NewAlert(&acorntypes.AlertParams{
		Content: "Alert",
	})

	alertOutlined := acorn.NewAlert(&acorntypes.AlertParams{
		Content: "Alert",
		Styles: &acorntypes.AlertStyles{
			Outlined: true,
		},
	})

	// NewButton
	buttonFilled := acorn.NewButton(&acorntypes.ButtonParams{
		Text: "Filled button",
		Link: "http://docs.thememountain.com/acorn/",
		Styles: &acorntypes.ButtonStyles{
			FullWidth: true,
		},
	})

	buttonOutlined := acorn.NewButton(&acorntypes.ButtonParams{
		Text: "Outlined button",
		Link: "http://docs.thememountain.com/acorn/",
		Styles: &acorntypes.ButtonStyles{
			Type: types.Outlined,
		},
	})

	buttonPhill := acorn.NewButton(&acorntypes.ButtonParams{
		Text: "Pill button",
		Link: "http://docs.thememountain.com/acorn/",
		Styles: &acorntypes.ButtonStyles{
			Type:  types.Pill,
			Align: aligns.Center,
		},
	})

	// NewAccordion
	accordion := acorn.NewAccordion([]*acorntypes.AccordionParams{
		{
			Title:   "Panel 1",
			Content: "Lorem ipsum dolor sit amet, consectetur adipisicing elit.",
		},
		{
			Title:   "Panel 2",
			Content: "Lorem ipsum dolor sit amet, consectetur adipisicing elit.",
			Styles: &acorntypes.AccordionStyles{
				Color:        colors.Cyan.M300,
				TitleColor:   colors.White,
				ContentColor: colors.Cyan.M300,
			},
		},
	})

	// NewSpacer
	spacer := acorn.NewSpacer()

	// NewDivider
	divider := acorn.NewDivider(colors.DeepPurple.M700)

	// Label
	labelFilled := acorn.NewLabel(&acorntypes.LabelParams{
		Text: "filled label",
	})

	labelOutlined := acorn.NewLabel(&acorntypes.LabelParams{
		Text: "outlined label",
		Styles: &acorntypes.LabelStyles{
			Outlined: true,
		},
	})

	// Content
	content := acorn.NewContent(&acorntypes.ContentParams{
		Content: row,
		Image:   "",
	})

	// withoutMargins := true
	// b := acorn.GetBoilerplate(acorntypes.AcornComponents{content}, &withoutMargins)

	timeline := acorn.NewTimeline([]*acorntypes.TimelineParams{
		{
			Time:    "2007 - 20082007 - 2008",
			Title:   "Massachusetts Institute of TechnologyMassachusetts Institute of TechnologyMassachusetts Institute of Technology",
			Content: "Co-wrote a book on the Semantic Web and Best Practices for Developers.Co-wrote a book on the Semantic Web and Best Practices for Developers.Co-wrote a book on the Semantic Web and Best Practices for Developers.",
		},
		{
			Time:    "2007 - 20082007 - 2008",
			Title:   "Massachusetts Institute of TechnologyMassachusetts Institute of TechnologyMassachusetts Institute of Technology",
			Content: "Co-wrote a book on the Semantic Web and Best Practices for Developers.Co-wrote a book on the Semantic Web and Best Practices for Developers.Co-wrote a book on the Semantic Web and Best Practices for Developers.",
		},
		{
			Time:    "2007 - 20082007 - 2008",
			Title:   "Massachusetts Institute of TechnologyMassachusetts Institute of TechnologyMassachusetts Institute of Technology",
			Content: "Co-wrote a book on the Semantic Web and Best Practices for Developers.Co-wrote a book on the Semantic Web and Best Practices for Developers.Co-wrote a book on the Semantic Web and Best Practices for Developers.",
		},
	})

	fmt.Println(timeline)

	boilertemplate := acorn.GetBoilerplate(acorntypes.AcornComponents{
		row,
		alert,
		buttonFilled,
		buttonOutlined,
		buttonPhill,
		accordion,
		spacer,
		divider,
		alertOutlined,
		labelFilled,
		labelOutlined,
		content,
		timeline,
	}, nil)

	t.Log(boilertemplate)
}

func TestGrid(t *testing.T) {

	acorn := acornmail.NewAcornEmailComponents()

	widthColumns := acornstyles.GetWidthColumns()
	colors := acornstyles.GetColors()

	grid := acorn.NewGrid([][]*acorntypes.ColumnParams{
		{
			{
				Content: "100%",
				Styles: &acorntypes.ColumnStyles{
					Width:     widthColumns.Full,
					Color:     colors.Purple.M700,
					TextColor: colors.White,
				},
			},
		},
		nil, // Add spacer
		{
			{
				Content: "1/2 de columna",
				Styles: &acorntypes.ColumnStyles{
					Width: widthColumns.Medium,
				},
			},
			{
				Content: "1/2 de columna",
				Styles: &acorntypes.ColumnStyles{
					Width: widthColumns.Medium,
				},
			},
		},
	})

	boilerplate := acorn.GetBoilerplate(acorntypes.AcornComponents{grid}, nil)

	t.Log(boilerplate)

}

func TestTestimonial(t *testing.T) {

	acorn := acornmail.NewAcornEmailComponents()

	aligns := acornstyles.GetAligns()
	colors := acornstyles.GetColors()

	testimonialBorder := acorn.NewTestimonial(&acorntypes.TestimonialParams{
		Testimonial: "Sometimes when you innovate, you make mistakes. It is best to admit them quickly, and get on with improving your other innovations.",
		Author:      "Steve Jobs",
		Styles: &acorntypes.TestimonialStyles{
			BorderColor: colors.Orange.M500,
		},
	})

	testimonialIcon := acorn.NewTestimonial(&acorntypes.TestimonialParams{
		Testimonial: "Sometimes when you innovate, you make mistakes. It is best to admit them quickly, and get on with improving your other innovations.",
		Author:      "Steve Jobs",
		Icon:        true,
	})

	testimonialImage := acorn.NewTestimonial(&acorntypes.TestimonialParams{
		Testimonial: "Sometimes when you innovate, you make mistakes. It is best to admit them quickly, and get on with improving your other innovations.",
		Author:      "Steve Jobs",
		Styles: &acorntypes.TestimonialStyles{
			Image: "https://gravatar.com/avatar/5ad269974f4c69c9ff6eca2ad2d1d0b8?s=400&d=robohash&r=x",
			Align: aligns.Center,
		},
	})

	boilerplate := acorn.GetBoilerplate(acorntypes.AcornComponents{
		testimonialBorder,
		testimonialIcon,
		testimonialImage,
	}, nil)

	t.Log(boilerplate)

}

func TestSpacer(t *testing.T) {
	acorn := acornmail.NewAcornEmailComponents()

	spacer := acorn.NewSpacer()

	boilerplate := acorn.GetBoilerplate(acorntypes.AcornComponents{spacer}, acornstyles.WithoutMargins())

	t.Log(boilerplate)
}

func TestDivider(t *testing.T) {
	acorn := acornmail.NewAcornEmailComponents()

	colors := acornstyles.GetColors()

	divider := acorn.NewDivider(colors.DeepPurple.M700)

	boilerplate := acorn.GetBoilerplate(acorntypes.AcornComponents{divider}, acornstyles.WithoutMargins())

	t.Log(boilerplate)
}

func TestRealExample(t *testing.T) {
	acorn := acornmail.NewAcornEmailComponents()

	aligns := acornstyles.GetAligns()
	colors := acornstyles.GetColors()

	variables := acorntypes.AcornVariables{
		"Name":  "Euclides Demóstenes",
		"Token": "Q2FwZXJ1Y2l0YSByb2ph",
	}

	image := acorn.NewImage(&acorntypes.ImageParams{
		Image: "https://i.picsum.photos/id/859/1200/280.jpg?hmac=cFup6pjvVaf67u1WSjrP8LYF8Oty0VrMKI3sbFDz8HQ",
		Alt:   "Logo",
	})

	button := acorn.NewButton(&acorntypes.ButtonParams{
		Text: "Activar cuenta",
		Link: "https://www.example.com?t={{.Token}}",
		Styles: &acorntypes.ButtonStyles{
			Align:     aligns.Center,
			Color:     colors.Cyan.M700,
			TextColor: colors.White,
		},
	})

	grid := acorn.NewGrid([][]*acorntypes.ColumnParams{
		{{Content: image}},
		{
			{
				Content: "<h1>¡Bienvenido!</h1>",
				Styles: &acorntypes.ColumnStyles{
					Align: aligns.Center,
				},
			},
		},
		{
			{
				Content: `
				Hola <b>{{.Name}}</b> gracias por registrarte en nuestro sitio web, para poder activar tu cuenta da click en el siguiente enlace:
				`,
				Styles: &acorntypes.ColumnStyles{
					Align: aligns.Center,
				},
			},
		},
		nil,
		{
			{
				Content: button,
				Styles: &acorntypes.ColumnStyles{
					Align: aligns.Center,
				},
			},
		},
	})

	boilerplate := acorn.GetBoilerplate(acorntypes.AcornComponents{grid}, nil)

	html, err := acornmail.MergeVariables(boilerplate, variables)
	if err != nil {
		panic(err)
	}

	t.Log(html)
}

func TestImage(t *testing.T) {
	acorn := acornmail.NewAcornEmailComponents()

	image := acorn.NewImage(&acorntypes.ImageParams{
		Image: "https://i.picsum.photos/id/859/1200/280.jpg?hmac=cFup6pjvVaf67u1WSjrP8LYF8Oty0VrMKI3sbFDz8HQ",
		Alt:   "Logo",
	})

	boilerplate := acorn.GetBoilerplate(acorntypes.AcornComponents{image}, nil)

	t.Log(boilerplate)
}

func TestPromo(t *testing.T) {
	acorn := acornmail.NewAcornEmailComponents()
	color := acornstyles.GetColors()
	aligns := acornstyles.GetAligns()

	promo := acorn.NewPromo(&acorntypes.PromoItems{
		Promo:       &acorntypes.PromoParams{Value: "25"},
		Symbol:      &acorntypes.PromoParams{Value: "%"},
		Description: &acorntypes.PromoParams{Value: "DESC"},
	})

	label := acorn.NewLabel(&acorntypes.LabelParams{
		Text: "25OFFTODAY",
		Styles: &acorntypes.LabelStyles{
			Outlined:  true,
			TextColor: color.Red.M500,
			Color:     color.Red.M500,
		},
	})

	coupon := acorn.NewCoupon(&acorntypes.CouponParams{
		Content: acorn.NewGrid([][]*acorntypes.ColumnParams{
			{{Content: promo}},
			{{
				Content: `Con el cupón ` + label,
				Styles: &acorntypes.ColumnStyles{
					Align:     aligns.Center,
					TextColor: color.Grey.M400,
				},
			}},
			nil,
		}),
		Button: &acorntypes.ButtonParams{
			Text: "Canjear cupón →",
			Styles: &acorntypes.ButtonStyles{
				Color:     color.Black,
				TextColor: color.White,
				Align:     aligns.Center,
			},
		},
		Styles: &acorntypes.CouponStyles{
			Dashed: true,
		},
	})

	fmt.Println(coupon)

	t.Log(coupon)
}

func TestCuopon(t *testing.T) {
	acorn := acornmail.NewAcornEmailComponents()
	color := acornstyles.GetColors()
	aligns := acornstyles.GetAligns()

	coupon := acorn.NewCoupon(&acorntypes.CouponParams{
		Content: acorn.NewGrid([][]*acorntypes.ColumnParams{
			{{
				Content: `
				<div style="font-size: 13px; text-transform: uppercase;">¡Gracias por registrarte!<br>Disfruta</div>
				<div style="font-size: 72px; font-weight: 700; line-height: 100%;">$10 DESC</div>
				<div class="spacer py-sm-8" style="line-height: 16px;">‌</div>
				<div style="font-size: 20px; letter-spacing: 2px; line-height: 100%; text-transform: uppercase;">En tu primera compra</div>
				`,
				Styles: &acorntypes.ColumnStyles{
					Align:     aligns.Center,
					TextColor: color.White,
				},
			}},
			nil,
		}),
		Button: &acorntypes.ButtonParams{
			Text: "COMPRA AHORA",
			Styles: &acorntypes.ButtonStyles{
				Color:     color.White,
				TextColor: color.Blue.M500,
				Align:     aligns.Center,
			},
		},
	})

	fmt.Println(coupon)

	t.Log(coupon)
}
