package acornstyles

import (
	"fmt"

	"github.com/arskang/gomail-acorn-template/acorntypes"
)

type sizes struct {
	Px12, Px24, Px36, Px48, Px60, Px72, Px84, Px96, Px108, Px120 *acorntypes.FontSize
}

func setSize(s int64) *acorntypes.FontSize {
	str := fmt.Sprint(s) + "px"
	return (*acorntypes.FontSize)(&str)
}

// Return sizes:
// Px12, Px24, Px36, Px48, Px60, Px72, Px84, Px96, Px108 and Px120
// (*acorntypes.FontSize)
func GetSizes() *sizes {
	return &sizes{
		Px12:  setSize(12),
		Px24:  setSize(24),
		Px36:  setSize(36),
		Px48:  setSize(48),
		Px60:  setSize(60),
		Px72:  setSize(72),
		Px84:  setSize(84),
		Px96:  setSize(96),
		Px108: setSize(108),
		Px120: setSize(120),
	}
}

// Create a new size
func NewAcornSize(s int64) *acorntypes.FontSize {
	return setSize(s)
}
