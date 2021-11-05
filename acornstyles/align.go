package acornstyles

import (
	"github.com/arskang/acornmail/acorntypes"
)

type aligns struct {
	Center, Right, Left *acorntypes.Align
}

func setAlign(a string) *acorntypes.Align {
	return (*acorntypes.Align)(&a)
}

// Return aligns:
// Center, Right and Left
// (*acorntypes.Align)
func GetAligns() *aligns {
	return &aligns{
		Center: setAlign("center"),
		Right:  setAlign("right"),
		Left:   setAlign("left"),
	}
}
