package acornstyles

import "github.com/arskang/gomail-acorn-template/acorntypes"

type widthColumns struct {
	Full, Quarter, Medium, ThreeQuarters, OneThird, TwoThird *acorntypes.WidthColumn
}

func setColumn(s string) *acorntypes.WidthColumn {
	return (*acorntypes.WidthColumn)(&s)
}

func GetWidthColumns() *widthColumns {
	return &widthColumns{
		Full:          setColumn("100%"),
		Quarter:       setColumn("138"),
		Medium:        setColumn("276"),
		ThreeQuarters: setColumn("414"),
		OneThird:      setColumn("184"),
		TwoThird:      setColumn("368"),
	}
}

func WithoutMargins() *bool {
	withoutMargins := true
	return &withoutMargins
}
