# Go Acorn Mail

#### Librerías
- [Acorn Email Framework](http://docs.thememountain.com/acorn/)

#### Instalación
```go
go get github.com/arskang/gomail-acorn-template
```

#### Ejemplo
```go
acorn := acornmail.NewAcornEmailComponents()
row := acorn.NewRow([]*acorntypes.ColParams{
    {
        Content: "Hola mundo",
    },
})
html := acorn.GetBoilerPlate(row)
fmt.Println(html)
```

#### Métodos
```go
import (
    acornmail "github.com/arskang/gomail-acorn-template"
)
```

- *GetHTMLString*: Obtener un html con parámetros dinámicos
```go
html, err := acornmail.GetHTMLString("<div>{{.Title}}</div>", map[string]interface{}{
    "Title": "Hola mundo",
})
if err != nil {
    panic(err)
}
fmt.Println(html)
```

- *NewAcornEmailComponents*: Obtener el tipo **acornEmail** para poder construir un template
```go
acorn := acornmail.NewAcornEmailComponents()
```

##### Components

- **GetBoilerPlate**
```go
boilertemplate := acorn.GetBoilerPlate(
    "Header",
    "Body",
    "Footer",
)
fmt.Println(boilertemplate)
```

- **Row**
```go
w := "1/4"
row := acorn.NewRow([]*acorntypes.ColParams{
    {
        Content: "Content",
        Width: &w,
    }
})
fmt.Println(row)
```

- **Alerts**
```go
content := "Aceptar"
hexColor := "#008f38"
outlined := nil // *bool
alert := acorn.NewAlert(content, hexColor, outlined)
fmt.Println(alert)
```

- **Buttons**
```go
button := acorn.NewButton(&acorntypes.ButtonParams {
    Text: "Aceptar",
    Link: "https://google.com",
    HexColor: "#008f38",
    // Align,
}, nil)
fmt.Println(button)
```

#### Types
```go
import "github.com/arskang/gomail-acorn-template/acorntypes"
```

- *Col*
```go
w := "1/4"
colParams := acorntypes.ColParams {
    Content: "Content", // string
    Width: &w, // *string
}
```

- *Button*
```go
buttonParams := acorntypes.ButtonParams {
    Text: "Aceptar",
    Link: "https://google.com",
    HexColor: "#008f38",
    // Align,
}
```

#### Utilidades

- *Ancho de columnas*

```go
// Grid a 4 columnas
w = "1/4"
w = "1/2"
w = "3/4"

// Grid a 3 columnas
w = "1/3"
w = "2/3"

// Todo el ancho
w = nil
w = "1"
w = "Cualquier otra cosa"

width := acorntypes.String(w)
```