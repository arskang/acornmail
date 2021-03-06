# Gomail Acorn Template

<!-- img src="https://d3vv6lp55qjaqc.cloudfront.net/items/1L1w0v431V0d1K410f3Y/keepAChangelog-logo-dark.svg" height=150 alt="Keep a Changelog" / -->

[![Go Report Card](https://goreportcard.com/badge/github.com/arskang/acornmail)](https://goreportcard.com/report/github.com/arskang/acornmail)
[![GoDoc](https://pkg.go.dev/badge/github.com/arskang/acornmail?status.svg)](https://pkg.go.dev/github.com/arskang/acornmail?tab=doc)
[![Release](https://img.shields.io/github/release/arskang/acornmail.svg?style=flat-square)](https://github.com/arskang/acornmail/releases)
[![MIT License Badge](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)
[![Keep a Changelog v1.1.0 badge](https://img.shields.io/badge/changelog-Keep%20a%20Changelog%20v1.1.0-%23E05735)](./CHANGELOG.md)

### Contenido

+ [Librerías](#librerías)
+ [Instalación](#instalación)
+ [Importación](#importación)
+ [Ejemplo](#ejemplo)
+ [Métodos](#métodos)
+ [Componentes](#componentes)
+ [Tipos](#tipos)
+ [Estilos](#estilos)
+ [Ejemplos reales](#ejemplos-reales)

#### Librerías
Proyecto original:
- [Acorn Email Framework](http://docs.thememountain.com/acorn/)

Colores:
- [Material design](https://material.io/resources/color/#!/?view.left=0&view.right=0)

#### Instalación
```
go get -u github.com/arskang/acornmail
```

#### Importación
```go
import "github.com/arskang/acornmail"
```

#### Ejemplo
Más ejemplos [aquí](./_examples/)
```go
import "github.com/arskang/acornmail"

func main() {

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

    fmt.Println(html)

}
```
![Ejemplo 01](./_assets/example-01.png)

#### Métodos

- **MergeVariables**: Fusionar a un HTML variables
```go
html, err := acornmail.MergeVariables(
    "<div>{{.Title}}</div>",
    acorntypes.AcornVariables{
        "Title": "Hola mundo",
    },
)
if err != nil {
    panic(err)
}

fmt.Println(html)
```

- **NewAcornEmailComponents**: Obtener el tipo **acornEmail** para poder construir un template
```go
acorn := acornmail.NewAcornEmailComponents()
```

#### Componentes

- **GetBoilerplate**
```go
boilerplate := acorn.GetBoilerplate(acorntypes.AcornComponents{
    "Header",
    "Body",
    "Footer",
    // n componentes...
}, nil)
fmt.Println(boilerplate)
```

- **Spacer**
```go
acorn := acornmail.NewAcornEmailComponents()

spacer := acorn.NewSpacer()

boilerplate := acorn.GetBoilerplate(
    acorntypes.AcornComponents{spacer},
    acornstyles.WithoutMargins(),
)

fmt.Println(boilerplate)
```
![Spacer](./_assets/components-spacer.png)

- **Divider**
```go
acorn := acornmail.NewAcornEmailComponents()

colors := acornstyles.GetColors()

divider := acorn.NewDivider(colors.DeepPurple.M700)

boilerplate := acorn.GetBoilerplate(
    acorntypes.AcornComponents{divider},
    acornstyles.WithoutMargins(),
)

fmt.Println(boilerplate)
```
![Divider](./_assets/components-divider.png)

- **Label**
```go
acorn := acornmail.NewAcornEmailComponents()

types := acornstyles.GetTypes()

labelFilled := acorn.NewLabel(&acorntypes.LabelParams{
    Text: "filled label",
})

labelOutlined := acorn.NewLabel(&acorntypes.LabelParams{
    Text: "outlined label",
    Styles: &acorntypes.LabelStyles{
        Outlined: true,
    },
})

boilerplate := acorn.GetBoilerplate(acorntypes.AcornComponents{
    labelFilled,
    labelOutlined,
}, nil)

fmt.Println(boilerplate)
```
![Label](./_assets/components-label.png)

- **Image**
```go
acorn := acornmail.NewAcornEmailComponents()

image := acorn.NewImage(&acorntypes.ImageParams{
    Image: "https://i.picsum.photos/id/859/1200/280.jpg?hmac=cFup6pjvVaf67u1WSjrP8LYF8Oty0VrMKI3sbFDz8HQ",
    Alt:   "Logo",
})

boilerplate := acorn.GetBoilerplate(acorntypes.AcornComponents{image}, nil)

fmt.Println(boilerplate)
```

![Label](./_assets/components-image.png)

- **Row**
```go
acorn := acornmail.NewAcornEmailComponents()

widthColumns := acornstyles.GetWidthColumns()

row := acorn.NewRow([]*acorntypes.ColumnParams{
    {
        Content: "1/4 de columna",
        Styles: &acorntypes.Styles{
            WidthColumn: widthColumns.Quarter,
        },
    },
    {
        Content: "1/2 de columna",
        Styles: &acorntypes.Styles{
            WidthColumn: widthColumns.Medium,
            Color:     colors.Purple.M700,
            TextColor: colors.White,
        },
    },
    {
        Content: "1/4 de columna",
        Styles: &acorntypes.Styles{
            WidthColumn: widthColumns.Quarter,
        },
    },
})

boilerplate := acorn.GetBoilerplate(acorntypes.AcornComponents{row}, nil)

fmt.Println(boilerplate)
```
![Row](./_assets/components-row.png)

- **Grid**
```go
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

fmt.Println(boilerplate)
```

![Grid](./_assets/components-grid.png)

- *Content*
```go
acorn := acornmail.NewAcornEmailComponents()

content := acorn.NewContent(&acorntypes.ContentParams{
    Content: row,
    Image:   "https://picsum.photos/1200/800?image=837",
})

boilerplate := acorn.GetBoilerplate(
    acorntypes.AcornComponents{content}, 
    acornstyles.WithoutMargins(),
)

fmt.Println(boilerplate)
```
![Content](./_assets/components-content.png)

```go
acorn := acornmail.NewAcornEmailComponents()

content := acorn.NewContent(&acorntypes.ContentParams{
    Content: row,
})

boilerplate := acorn.GetBoilerplate(
    acorntypes.AcornComponents{content}, 
    acornstyles.WithoutMargins(),
)

fmt.Println(boilerplate)
```
![Content](./_assets/components-content-withoutimage.png)

- **Alerts**
```go
acorn := acornmail.NewAcornEmailComponents()

divider := acorn.NewDivider(nil)

alert := acorn.NewAlert(&acorntypes.AlertParams{
    Content: "Alert",
})

alertOutlined := acorn.NewAlert(&acorntypes.AlertParams{
    Content: "Alert",
    Styles: &acorntypes.AlertStyles{
        Outlined: true,
    },
})

boilerplate := acorn.GetBoilerplate(acorntypes.AcornComponents{
    alert,
    divider,
    alertOutlined,
}, nil)

fmt.Println(boilerplate)
```
![Alert](./_assets/components-alert.png)

- **Buttons**
```go
acorn := acornmail.NewAcornEmailComponents()

widthColumns := acornstyles.GetWidthColumns()
types := acornstyles.GetTypes()
aligns := acornstyles.GetAligns()

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

grid := acorn.NewGrid([][]*acorntypes.ColumnParams{
    {
        {
            Content: buttonFilled,
            Styles: &acorntypes.ColumnStyles{
                Width: widthColumns.Full,
            },
        },
    },
    nil, // Add spacer
    {
       {
            Content: buttonOutlined,
            Styles: &acorntypes.ColumnStyles{
                Width: widthColumns.Medium,
            },
        },
        {
            Content: buttonPhill,
            Styles: &acorntypes.ColumnStyles{
                Width: widthColumns.Medium,
            },
        },
    },
})

boilerplate := acorn.GetBoilerplate(acorntypes.AcornComponents{grid}, nil)

fmt.Println(boilerplate)
```
![Buttons](./_assets/components-buttons.png)

- **Accordion**
```go
acorn := acornmail.NewAcornEmailComponents()

colors := acornstyles.GetColors()

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

boilerplate := acorn.GetBoilerplate(acorntypes.AcornComponents{accordion}, nil)

fmt.Println(boilerplate)
```
![Accordion](./_assets/components-accordion.png)

- **Timeline**
```go
acorn := acornmail.NewAcornEmailComponents()

//== Limitaciones ==//
// Time: Max 11 caracteres
// Title: Max 37 caracteres
// Content: Max 78 caracteres

timeline := acorn.NewTimeline([]*acorntypes.TimelineParams{
    {
        Time:    "2007 - 2008",
        Title:   "Massachusetts Institute of Technology",
        Content: "Co-wrote a book on the Semantic Web and Best Practices for Developers.",
    },
    {
        Time:    "2004 - 2006",
        Title:   "Parsons School of Design",
        Content: "Awarded Best Designer of the Year voted by student and teacher body.",
    },
    {
        Time:    "2002 - 2004",
        Title:   "Berkley College",
        Content: "Two year program with a focus on design principles and Javascript development.",
    },
})

boilerplate := acorn.GetBoilerplate(acorntypes.AcornComponents{timeline}, nil)

fmt.Println(boilerplate)
```
![Timeline](./_assets/components-timeline.png)

- **Testimonial**
```go
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

fmt.Println(boilerplate)
```
![Testimonial](./_assets/components-testimonial.png)

- **Promo**
```go
acorn := acornmail.NewAcornEmailComponents()

promo := acorn.NewPromo(&acorntypes.PromoItems{
    Promo: &acorntypes.PromoParams{Value: "25"},
    Symbol: &acorntypes.PromoParams{Value: "%"},
    Description: &acorntypes.PromoParams{Value: "OFF"},
})

boilerplate := acorn.GetBoilerplate(acorntypes.AcornComponents{
    promo,
}, nil)
```
![Promo default](./_assets/components-promo01.png)

```go
acorn := acornmail.NewAcornEmailComponents()
sizes := acornstyles.GetSizes()

promo := acorn.NewPromo(&acorntypes.PromoItems{
    Promo: &acorntypes.PromoParams{
        Value: "25",
        Size:  sizes.PX108,
    },
    Symbol: &acorntypes.PromoParams{
        Value: "%",
        Color: acornstyles.GetColors().Green.M700,
        Size:  sizes.PX36,
    },
    Description: &acorntypes.PromoParams{
        Value: "OFF",
        Size:  sizes.PX24,
    },
})

boilerplate := acorn.GetBoilerplate(acorntypes.AcornComponents{
    promo,
}, nil)
```
![Promo custom](./_assets/components-promo02.png)

- **Coupon**
```go
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

boilerplate := acorn.GetBoilerplate(acorntypes.AcornComponents{
    coupon,
}, acornstyles.WithoutMargins())
```
![Coupon default](./_assets/components-coupon01.png)

```go
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

boilerplate := acorn.GetBoilerplate(acorntypes.AcornComponents{
    coupon,
}, acornstyles.WithoutMargins())
```
![Coupon dashed](./_assets/components-coupon02.png)

#### Tipos

- *Importación*
```go
import "github.com/arskang/acornmail/acorntypes"
```

- *Básicos*

    - Align ```acorntypes.Align```
    - Color ```acorntypes.Color```
    - WidthColumn ```acorntypes.WidthColumn```
    - Types ```acorntypes.Types```
    - AcornComponents ```acorntypes.AcornComponents```
    - AcornVariables ```acorntypes.AcornVariables```
    - FontSize ```acorntypes.FontSize```

- *Compuestos*

    - ColumnParams ```acorntypes.ColumnParams```
    - ColumnStyles ```acorntypes.ColumnStyles```
    - ButtonParams ```acorntypes.ButtonParams```
    - ButtonStyles ```acorntypes.ButtonStyles```
    - AlertParams ```acorntypes.AlertParams```
    - AlertStyles ```acorntypes.AlertStyles```
    - AccordionParams ```acorntypes.AccordionParams```
    - AccordionStyles ```acorntypes.AccordionStyles```
    - LabelParams ```acorntypes.LabelParams```
    - LabelStyles ```acorntypes.LabelStyles```
    - ContentParams ```acorntypes.ContentParams```
    - TimelineParams ```acorntypes.TimelineParams```
    - TestimonialParams ```acorntypes.TestimonialParams```
    - TestimonialStyles ```acorntypes.TestimonialStyles```
    - ImageParams ```acorntypes.ImageParams```
    - CouponParams ```acorntypes.CouponParams```
    - PromoItems ```acorntypes.PromoItems```
    - PromoParams ```acorntypes.PromoParams```
    - CouponStyles ```acorntypes.CouponStyles```

#### Estilos

- *Importación*
```go
import "github.com/arskang/acornmail/acornstyles"
```

- *Tipo de botones*

```go
types := acornstyles.GetTypes()
fmt.Println(types.Filled)
fmt.Println(types.Outlined)
fmt.Println(types.Pill)
```

- *Alineaciones*

```go
aligns := acornstyles.GetAligns()
fmt.Println(aligns.Center)
fmt.Println(aligns.Right)
fmt.Println(aligns.Left)
```

- *Tamaños de letra*

```go
sizes := acornstyles.GetSizes()
fmt.Println(sizes.Px12)
fmt.Println(sizes.Px24)
fmt.Println(sizes.Px36)
fmt.Println(sizes.Px48)
fmt.Println(sizes.Px60)
fmt.Println(sizes.Px72)
fmt.Println(sizes.Px84)
fmt.Println(sizes.Px96)
fmt.Println(sizes.Px108)
fmt.Println(sizes.Px120)

customSize := acornstyles.NewAcornSize(10)
fmt.Println(customSize)
// 10px
```

- *Ancho de columnas*

```go
widthColumn := acornstyles.GetWidthColumns()
fmt.Println(widthColumn.Full) // 100%
fmt.Println(widthColumn.Quarter) // 1/4
fmt.Println(widthColumn.Medium) // 1/2
fmt.Println(widthColumn.ThreeQuarters) // 3/4
fmt.Println(widthColumn.OneThird) // 1/3
fmt.Println(widthColumn.TwoThird) // 2/3
```

![Grid three](./_assets/grid-three.png)
![Grid four](./_assets/grid-four.png)

- *Colores*

```go
colors := acornstyles.GetColors()
red50 := colors.Red.M50
fmt.Println(red50)

// Only hexadecimal string
customColor, err := acornstyles.NewAcornColor("#fea800")
if err != nil {
    panic(err)
}
fmt.Println(customColor)
```

![Material color 01](./_assets/material-color-01.jpeg)
![Material color 02](./_assets/material-color-02.jpeg)
![Material color 03](./_assets/material-color-03.jpeg)
