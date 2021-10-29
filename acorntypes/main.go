package acorntypes

// Type Align (string)
type Align string

// Return a string
func (a Align) String() string {
	return string(a)
}

// Type Color (string)
type Color string

// Return a string
func (c Color) String() string {
	return string(c)
}

// Type WidthColumn (string)
type WidthColumn string

// Return a string
func (wc WidthColumn) String() string {
	return string(wc)
}

// Type Types (string)
type Types string

// Return a string
func (t Types) String() string {
	return string(t)
}

// Type AcornComponents ([]string)
type AcornComponents []string

// Type AcornVariables (map[string]string)
type AcornVariables map[string]string

// Type ColumnParams
type ColumnParams struct {
	Content string
	Styles  *ColumnStyles
}

// Type ButtonParams
type ButtonParams struct {
	Text   string
	Link   string
	Styles *ButtonStyles
}

// Type AlertParams
type AlertParams struct {
	Content string
	Styles  *AlertStyles
}

// Type AccordionParams
type AccordionParams struct {
	Title   string
	Content string
	Styles  *AccordionStyles
}

// Type LabelParams
type LabelParams struct {
	Text   string
	Styles *LabelStyles
}

// Type ContentParams
type ContentParams struct {
	Content string
	Image   string
	Color   *Color
}

// Type TimelineParams
type TimelineParams struct {
	Time, Title, Content string
}

// Type TestimonialParams
type TestimonialParams struct {
	Testimonial string
	Author      string
	Icon        bool
	Styles      *TestimonialStyles
}

// Type ImageParams
type ImageParams struct {
	Image       string
	Alt         string
	WidthColumn *WidthColumn
}

// Type ColumnStyles
type ColumnStyles struct {
	Align     *Align
	Width     *WidthColumn
	Color     *Color
	TextColor *Color
}

// Type ButtonStyles
type ButtonStyles struct {
	FullWidth bool
	Align     *Align
	Color     *Color
	TextColor *Color
	Type      *Types
}

// Type AlertStyles
type AlertStyles struct {
	Outlined  bool
	Color     *Color
	TextColor *Color
}

// Type AccordionStyles
type AccordionStyles struct {
	Color        *Color
	TitleColor   *Color
	ContentColor *Color
}

// Type LabelStyles
type LabelStyles struct {
	Outlined  bool
	Color     *Color
	TextColor *Color
	Type      *Types // Deprecated: Uses better Outlined
}

// Type TestimonialStyles
type TestimonialStyles struct {
	Image       string
	Align       *Align
	BorderColor *Color
}
