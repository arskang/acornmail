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

type AcornComponents []string

type AcornVariables map[string]string

type ColumnParams struct {
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

type AccordionParams struct {
	Title   string
	Content string
	Styles  *AccordionStyles
}

type LabelParams struct {
	Text   string
	Styles *LabelStyles
}

type ContentParams struct {
	Content string
	Image   string
	Color   *Color
}

type ColumnStyles struct {
	Align     *Align
	Width     *WidthColumn
	Color     *Color
	TextColor *Color
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

type AccordionStyles struct {
	Color        *Color
	TitleColor   *Color
	ContentColor *Color
}

type LabelStyles struct {
	Type      *Types
	Color     *Color
	TextColor *Color
}
