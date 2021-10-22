package acornstyles

import (
	"github.com/arskang/gomail-acorn-template/acorntypes"
)

type aligns struct {
	Center, Right, Left *acorntypes.Align
}

func setAlign(a string) *acorntypes.Align {
	return (*acorntypes.Align)(&a)
}

func GetAligns() *aligns {
	return &aligns{
		Center: setAlign("center"),
		Right:  setAlign("right"),
		Left:   setAlign("left"),
	}
}
