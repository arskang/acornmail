package acorntypes

type Align string

func (a Align) String() string {
	return string(a)
}

type Color string

func (c Color) String() string {
	return string(c)
}

type WidthColumn string

func (wc WidthColumn) String() string {
	return string(wc)
}

type Types string

func (t Types) String() string {
	return string(t)
}

type AcornVariables map[string]string

type ColParams struct {
	Content string
	Styles  *ColumnStyles
}

type ButtonParams struct {
	Text   string
	Link   string
	Styles *ButtonStyles
}

type AlertParams struct {
	Content string
	Styles  *AlertStyles
}

type ColumnStyles struct {
	Align *Align
	Width *WidthColumn
}

type ButtonStyles struct {
	FullWidth bool
	Align     *Align
	Color     *Color
	TextColor *Color
	Type      *Types
}

type AlertStyles struct {
	Outlined  bool
	Color     *Color
	TextColor *Color
}
