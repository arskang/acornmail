package acornstyles

import (
	"fmt"
	"regexp"

	"github.com/arskang/gomail-acorn-template/acorntypes"
)

type aligns struct {
	Center, Right, Left acorntypes.Align
}

type widthColumns struct {
	Full, Quarter, Medium, ThreeQuarters, OneThird, TwiThird *acorntypes.WidthColumn
}

func GetAligns() *aligns {
	return &aligns{
		Center: "center",
		Right:  "right",
		Left:   "left",
	}
}

func setColumn(s string) *acorntypes.WidthColumn {
	return (*acorntypes.WidthColumn)(&s)
}

func setColor(s string) *acorntypes.Color {
	return (*acorntypes.Color)(&s)
}

func GetColor(hex string) (*acorntypes.Color, error) {
	match, _ := regexp.MatchString(RGX_HEXCOLOR, hex)
	if match {
		return setColor(hex), nil
	}
	return nil, fmt.Errorf("it is not a hexadecimal color")
}

func GetWidthColumns() *widthColumns {
	return &widthColumns{
		Full:          setColumn("100%"),
		Quarter:       setColumn("138"),
		Medium:        setColumn("276"),
		ThreeQuarters: setColumn("414"),
		OneThird:      setColumn("184"),
		TwiThird:      setColumn("368"),
	}
}
