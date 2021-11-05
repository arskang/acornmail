package acornstyles

import "github.com/arskang/acornmail/acorntypes"

type types struct {
	Filled, Outlined, Pill *acorntypes.Types
}

func setType(t string) *acorntypes.Types {
	return (*acorntypes.Types)(&t)
}

// Return types:
// Filled, Outlined and Pill (*acorntypes.Types)
func GetTypes() *types {
	return &types{
		Filled:   setType("filled"),
		Outlined: setType("outlined"),
		Pill:     setType("pill"),
	}
}
