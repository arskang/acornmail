# Gomail Acorn Template

+ [Librerías](#librerías)
+ [Instalación](#instalación)
+ [Importación](#importación)
+ [Ejemplo](#ejemplo)
+ [Métodos](#métodos)
+ [Componentes](#componentes)
+ [Tipos](#tipos)
+ [Estilos](#estilos)

#### Librerías
Proyecto original:
- [Acorn Email Framework](http://docs.thememountain.com/acorn/)

Colores:
- [Material design](https://material.io/resources/color/#!/?view.left=0&view.right=0)

#### Instalación
```
go get -u github.com/arskang/gomail-acorn-template
```

#### Importación
```go
import acornmail "github.com/arskang/gomail-acorn-template"
```

#### Ejemplo
```go
import acornmail "github.com/arskang/gomail-acorn-template"

func main() {

    acorn := acornmail.NewAcornEmailComponents()

}

```

#### Métodos

- *MergeVariables*: Fusionar a un HTML variables
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

- *NewAcornEmailComponents*: Obtener el tipo **acornEmail** para poder construir un template
```go
acorn := acornmail.NewAcornEmailComponents()
```

##### Componentes

- **GetBoilerPlate**
```go
boilertemplate := acorn.GetBoilerPlate(
    "Header",
    "Body",
    "Footer",
    // n componentes...
)
fmt.Println(boilertemplate)
```

- **Row**
```go
wColumns := acornstyles.GetWidthColumns()

row := acorn.NewRow([]*acorntypes.ColParams{
    {
        Content: "1/4 de columna",
        Styles: &acorntypes.Styles{
            WidthColumn: wColumns.Quarter,
        },
    },
    {
        Content: "1/2 de columna",
        Styles: &acorntypes.Styles{
            WidthColumn: wColumns.Medium,
        },
    },
    {
        Content: "1/4 de columna",
        Styles: &acorntypes.Styles{
            WidthColumn: wColumns.Quarter,
        },
    },
})

fmt.Println(row)
```

- **Alerts**
```go
colors := acornstyles.GetColors()

alert := acorn.NewAlert(&acorntypes.AlertParams{
    Content: "Esta es una alerta",
    Styles: &acorntypes.AlertStyles{
        Outlined:  true,
        TextColor: colors.Pink.M300,
    },
})

fmt.Println(alert)
```

- **Buttons**
```go
types := acornstyles.GetTypes()
aligns := acornstyles.GetAligns()

buttonFilled := acorn.NewButton(&acorntypes.ButtonParams{
    Text: "Filled button",
    Link: "http://docs.thememountain.com/acorn/",
    Styles: &acorntypes.ButtonStyles{
        FullWidth: true,
    },
})

fmt.Println(buttonFilled)

buttonOutlined := acorn.NewButton(&acorntypes.ButtonParams{
    Text: "Outlined button",
    Link: "http://docs.thememountain.com/acorn/",
    Styles: &acorntypes.ButtonStyles{
        Type: types.Outlined,
    },
})

fmt.Println(buttonOutlined)

buttonPhill := acorn.NewButton(&acorntypes.ButtonParams{
    Text: "Pill button",
    Link: "http://docs.thememountain.com/acorn/",
    Styles: &acorntypes.ButtonStyles{
        Type:  types.Pill,
        Align: aligns.Center,
    },
})

fmt.Println(buttonPhill)
```

#### Tipos

- *Importación*
```go
import "github.com/arskang/gomail-acorn-template/acorntypes"
```

- *Básicos*

    - Align ```acorntypes.Align```
    - Color ```acorntypes.Color```
    - WidthColumn ```acorntypes.WidthColumn```
    - Types ```acorntypes.Types```
    - AcornVariables ```acorntypes.AcornVariables```

- *Compuestos*

    - ColParams ```acorntypes.ColParams```
    - ButtonParams ```acorntypes.ButtonParams```
    - AlertParams ```acorntypes.AlertParams```
    - ColumnStyles ```acorntypes.ColumnStyles```
    - ButtonStyles ```acorntypes.ButtonStyles```
    - AlertStyles ```acorntypes.AlertStyles```

#### Estilos

- *Importación*
```go
import "github.com/arskang/gomail-acorn-template/acornstyles"
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

![Grid three](./assets/grid-three.png)
![Grid four](./assets/grid-four.png)

- *Colores*

```go
acornColor := acornstyles.GetColors()
red50 := acornColor.Red.M50
fmt.Println(red50)

// Only hexadecimal string
customColor, err := acornstyles.NewAcornColor("#fea800")
if err != nil {
    panic(err)
}
fmt.Println(customColor)
```

![Material color 01](./assets/material-color-01.jpeg)
![Material color 02](./assets/material-color-02.jpeg)
![Material color 03](./assets/material-color-03.jpeg)