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

// Type FontSize (string)
type FontSize string

// Return a string
func (fs FontSize) String() string {
	return string(fs)
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

// Type CouponParams
type CouponParams struct {
	Content string
	Button  *ButtonParams
	Styles  *CouponStyles
}

// Type PromoItems
type PromoItems struct {
	Promo       *PromoParams
	Symbol      *PromoParams
	Description *PromoParams
}

// Type PromoParams
type PromoParams struct {
	Value string
	Color *Color
	Size  *FontSize
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
}

// Type TestimonialStyles
type TestimonialStyles struct {
	Image       string
	Align       *Align
	BorderColor *Color
}

// Type CouponStyles
type CouponStyles struct {
	Dashed      bool
	Color       *Color
	BorderColor *Color
}
