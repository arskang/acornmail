package acornstyles

import "github.com/arskang/acornmail/acorntypes"

type widthColumns struct {
	Full, Quarter, Medium, ThreeQuarters, OneThird, TwoThird *acorntypes.WidthColumn
}

func setColumn(s string) *acorntypes.WidthColumn {
	return (*acorntypes.WidthColumn)(&s)
}

// Return width columns:
// Full, Quarter, Medium, ThreeQuarters, OneThird and TwoThird
// (*acorntypes.WidthColumn)
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
