package acornstyles

import "github.com/arskang/gomail-acorn-template/acorntypes"

type colorPalette struct {
	_50, _100, _200, _300, _400, _500, _600, _700, _800, _900 acorntypes.Color
}

type colorPaletteA struct {
	_50, _100, _200, _300, _400, _500, _600, _700, _800, _900, A100, A200, A400, A700 acorntypes.Color
}

type colors struct {
	White      acorntypes.Color
	Black      acorntypes.Color
	Brown      colorPalette
	Grey       colorPalette
	BlueGrey   colorPalette
	Red        colorPaletteA
	Pink       colorPaletteA
	Purple     colorPaletteA
	DeepPurple colorPaletteA
	Indigo     colorPaletteA
	Blue       colorPaletteA
	LightBlue  colorPaletteA
	Cyan       colorPaletteA
	Teal       colorPaletteA
	Green      colorPaletteA
	LightGreen colorPaletteA
	Lime       colorPaletteA
	Yellow     colorPaletteA
	Amber      colorPaletteA
	Orange     colorPaletteA
	DeepOrange colorPaletteA
}

func GetColors() *colors {
	return &colors{
		White: "#ffffff",
		Black: "#21292f",
		Brown: colorPalette{
			_50:  "#efebe9",
			_100: "#d7ccc8",
			_200: "#bcaaa4",
			_300: "#a1887f",
			_400: "#8d6e63",
			_500: "#795548",
			_600: "#6d4c41",
			_700: "#5d4037",
			_800: "#4e342e",
			_900: "#3e2723",
		},
		Grey:     colorPalette{},
		BlueGrey: colorPalette{},
		Red: colorPaletteA{
			_50:  "#ffebee",
			_100: "#ffcdd2",
			_200: "#ef9a9a",
			_300: "#e57373",
			_400: "#ef5350",
			_500: "#f44336",
			_600: "#e53935",
			_700: "#d32f2f",
			_800: "#c62828",
			_900: "#b71c1c",
			A100: "#ff5252",
			A200: "#ff5252",
			A400: "#ff1744",
			A700: "#d50000",
		},
		Pink:       colorPaletteA{},
		Purple:     colorPaletteA{},
		DeepPurple: colorPaletteA{},
		Indigo:     colorPaletteA{},
		Blue:       colorPaletteA{},
		LightBlue:  colorPaletteA{},
		Cyan:       colorPaletteA{},
		Teal:       colorPaletteA{},
		Green:      colorPaletteA{},
		LightGreen: colorPaletteA{},
		Lime:       colorPaletteA{},
		Yellow:     colorPaletteA{},
		Amber:      colorPaletteA{},
		Orange:     colorPaletteA{},
		DeepOrange: colorPaletteA{},
	}
}
