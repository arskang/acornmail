package acorntypes

type Align string

type Color string

type WidthColumn string

type ColParams struct {
	Content string
	Styles  *ColumnStyles
}

type ButtonParams struct {
	Text, Link, HexColor, Align string
}

type Styles struct {
	Align       *Align
	Color       *Color
	WidthColumn *WidthColumn
}

type ColumnStyles struct {
	Align *Align
	Width *WidthColumn
}

type ButtonStyles struct {
	Align *Align
	Color *Color
}

func String(s string) *string {
	return &s
}
