package acornstyles

import "github.com/arskang/gomail-acorn-template/acorntypes"

type types struct {
	Filled, Outlined, Pill *acorntypes.Types
}

func setType(t string) *acorntypes.Types {
	return (*acorntypes.Types)(&t)
}

func GetTypes() *types {
	return &types{
		Filled:   setType("filled"),
		Outlined: setType("outlined"),
		Pill:     setType("pill"),
	}
}
