package acornstyles

import (
	"github.com/arskang/gomail-acorn-template/acorntypes"
)

type aligns struct {
	Center, Right, Left acorntypes.Align
}

func GetAligns() *aligns {
	return &aligns{
		Center: "center",
		Right:  "right",
		Left:   "left",
	}
}
