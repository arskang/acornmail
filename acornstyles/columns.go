package acornstyles

import (
	"fmt"

	"github.com/arskang/acornmail/acorntypes"
)

type widthColumns struct {
	Full, Quarter, Medium, ThreeQuarters, OneThird, TwoThird *acorntypes.WidthColumn
}

func setColumn(s string) *acorntypes.WidthColumn {
	return (*acorntypes.WidthColumn)(&s)
}

func calculateWidths(width int) map[string]int {
	return map[string]int{
		"Quarter":       width / 4,
		"Medium":        width / 2,
		"ThreeQuarters": (3 * width) / 4,
		"OneThird":      width / 3,
		"TwoThird":      (2 * width) / 3,
	}
}

// Return width columns:
// Full, Quarter, Medium, ThreeQuarters, OneThird and TwoThird
// (*acorntypes.WidthColumn)
func GetWidthColumns(width ...int) *widthColumns {
	defaultWidth := 600
	if len(width) > 0 {
		defaultWidth = width[0]
	}
	values := calculateWidths(defaultWidth)
	WidthColumns := &widthColumns{
		Full: setColumn("100%"),
	}
	for key, v := range values {
		switch key {
		case "Quarter":
			WidthColumns.Quarter = setColumn(fmt.Sprint(v))
		case "Medium":
			WidthColumns.Medium = setColumn(fmt.Sprint(v))
		case "ThreeQuarters":
			WidthColumns.ThreeQuarters = setColumn(fmt.Sprint(v))
		case "OneThird":
			WidthColumns.OneThird = setColumn(fmt.Sprint(v))
		case "TwoThird":
			WidthColumns.TwoThird = setColumn(fmt.Sprint(v))
		}
	}
	return WidthColumns
}

func WithoutMargins() *bool {
	withoutMargins := true
	return &withoutMargins
}
