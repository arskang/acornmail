package acornstyles

import (
	"fmt"
	"regexp"

	"github.com/arskang/gomail-acorn-template/acorntypes"
)

type colorPalette struct {
	M50, M100, M200, M300, M400, M500, M600, M700, M800, M900 *acorntypes.Color
}

type colorPaletteA struct {
	M50, M100, M200, M300, M400, M500, M600, M700, M800, M900, A100, A200, A400, A700 *acorntypes.Color
}

type colors struct {
	White      *acorntypes.Color
	Black      *acorntypes.Color
	Brown      *colorPalette
	Grey       *colorPalette
	BlueGrey   *colorPalette
	Red        *colorPaletteA
	Pink       *colorPaletteA
	Purple     *colorPaletteA
	DeepPurple *colorPaletteA
	Indigo     *colorPaletteA
	Blue       *colorPaletteA
	LightBlue  *colorPaletteA
	Cyan       *colorPaletteA
	Teal       *colorPaletteA
	Green      *colorPaletteA
	LightGreen *colorPaletteA
	Lime       *colorPaletteA
	Yellow     *colorPaletteA
	Amber      *colorPaletteA
	Orange     *colorPaletteA
	DeepOrange *colorPaletteA
}

func setColor(s string) *acorntypes.Color {
	return (*acorntypes.Color)(&s)
}

func IsHexColor(hex string) bool {
	match, err := regexp.MatchString(RGX_HEXCOLOR, hex)
	if err != nil {
		return false
	}
	return match
}

func NewAcornColor(hex string) (*acorntypes.Color, error) {
	if IsHexColor(hex) {
		return setColor(hex), nil
	}
	return nil, fmt.Errorf("it is not a hexadecimal color")
}

func GetColors() *colors {
	return &colors{
		White: setColor("#ffffff"),
		Black: setColor("#000000"),
		Brown: &colorPalette{
			M50:  setColor("#efebe9"),
			M100: setColor("#d7ccc8"),
			M200: setColor("#bcaaa4"),
			M300: setColor("#a1887f"),
			M400: setColor("#8d6e63"),
			M500: setColor("#795548"),
			M600: setColor("#6d4c41"),
			M700: setColor("#5d4037"),
			M800: setColor("#4e342e"),
			M900: setColor("#3e2723"),
		},
		Grey: &colorPalette{
			M50:  setColor("#fafafa"),
			M100: setColor("#d7ccc8"),
			M200: setColor("#eeeeee"),
			M300: setColor("#e0e0e0"),
			M400: setColor("#bdbdbd"),
			M500: setColor("#9e9e9e"),
			M600: setColor("#757575"),
			M700: setColor("#616161"),
			M800: setColor("#424242"),
			M900: setColor("#212121"),
		},
		BlueGrey: &colorPalette{
			M50:  setColor("#eceff1"),
			M100: setColor("#cfd8dc"),
			M200: setColor("#b0bec5"),
			M300: setColor("#90a4ae"),
			M400: setColor("#78909c"),
			M500: setColor("#607d8b"),
			M600: setColor("#546e7a"),
			M700: setColor("#455a64"),
			M800: setColor("#37474f"),
			M900: setColor("#263238"),
		},
		Red: &colorPaletteA{
			M50:  setColor("#ffebee"),
			M100: setColor("#ffcdd2"),
			M200: setColor("#ef9a9a"),
			M300: setColor("#e57373"),
			M400: setColor("#ef5350"),
			M500: setColor("#f44336"),
			M600: setColor("#e53935"),
			M700: setColor("#d32f2f"),
			M800: setColor("#c62828"),
			M900: setColor("#b71c1c"),
			A100: setColor("#ff5252"),
			A200: setColor("#ff5252"),
			A400: setColor("#ff1744"),
			A700: setColor("#d50000"),
		},
		Pink: &colorPaletteA{
			M50:  setColor("#fce4ec"),
			M100: setColor("#f8bbd0"),
			M200: setColor("#f48fb1"),
			M300: setColor("#f06292"),
			M400: setColor("#ec407a"),
			M500: setColor("#e91e63"),
			M600: setColor("#d81b60"),
			M700: setColor("#c2185b"),
			M800: setColor("#ad1457"),
			M900: setColor("#880e4f"),
			A100: setColor("#ff80ab"),
			A200: setColor("#ff4081"),
			A400: setColor("#f50057"),
			A700: setColor("#c51162"),
		},
		Purple: &colorPaletteA{
			M50:  setColor("#f3e5f5"),
			M100: setColor("#e1bee7"),
			M200: setColor("#ce93d8"),
			M300: setColor("#ba68c8"),
			M400: setColor("#ab47bc"),
			M500: setColor("#9c27b0"),
			M600: setColor("#8e24aa"),
			M700: setColor("#7b1fa2"),
			M800: setColor("#6a1b9a"),
			M900: setColor("#4a148c"),
			A100: setColor("#ea80fc"),
			A200: setColor("#e040fb"),
			A400: setColor("#d500f9"),
			A700: setColor("#aa00ff"),
		},
		DeepPurple: &colorPaletteA{
			M50:  setColor("#ede7f6"),
			M100: setColor("#d1c4e9"),
			M200: setColor("#b39ddb"),
			M300: setColor("#9575cd"),
			M400: setColor("#7e57c2"),
			M500: setColor("#673ab7"),
			M600: setColor("#5e35b1"),
			M700: setColor("#512da8"),
			M800: setColor("#4527a0"),
			M900: setColor("#311b92"),
			A100: setColor("#b388ff"),
			A200: setColor("#7c4dff"),
			A400: setColor("#651fff"),
			A700: setColor("#6200ea"),
		},
		Indigo: &colorPaletteA{
			M50:  setColor("#e8eaf6"),
			M100: setColor("#c5cae9"),
			M200: setColor("#9fa8da"),
			M300: setColor("#7986cb"),
			M400: setColor("#5c6bc0"),
			M500: setColor("#3f51b5"),
			M600: setColor("#3949ab"),
			M700: setColor("#303f9f"),
			M800: setColor("#283593"),
			M900: setColor("#1a237e"),
			A100: setColor("#8c9eff"),
			A200: setColor("#536dfe"),
			A400: setColor("#3d5afe"),
			A700: setColor("#304ffe"),
		},
		Blue: &colorPaletteA{
			M50:  setColor("#e3f2fd"),
			M100: setColor("#bbdefb"),
			M200: setColor("#90caf9"),
			M300: setColor("#64b5f6"),
			M400: setColor("#42a5f5"),
			M500: setColor("#2196f3"),
			M600: setColor("#1e88e5"),
			M700: setColor("#1976d2"),
			M800: setColor("#1565c0"),
			M900: setColor("#0d47a1"),
			A100: setColor("#82b1ff"),
			A200: setColor("#448aff"),
			A400: setColor("#2979ff"),
			A700: setColor("#2962ff"),
		},
		LightBlue: &colorPaletteA{
			M50:  setColor("#e1f5fe"),
			M100: setColor("#b3e5fc"),
			M200: setColor("#81d4fa"),
			M300: setColor("#4fc3f7"),
			M400: setColor("#29b6f6"),
			M500: setColor("#03a9f4"),
			M600: setColor("#039be5"),
			M700: setColor("#0288d1"),
			M800: setColor("#0277bd"),
			M900: setColor("#01579b"),
			A100: setColor("#80d8ff"),
			A200: setColor("#40c4ff"),
			A400: setColor("#00b0ff"),
			A700: setColor("#0091ea"),
		},
		Cyan: &colorPaletteA{
			M50:  setColor("#e0f7fa"),
			M100: setColor("#b2ebf2"),
			M200: setColor("#80deea"),
			M300: setColor("#4dd0e1"),
			M400: setColor("#26c6da"),
			M500: setColor("#00bcd4"),
			M600: setColor("#00acc1"),
			M700: setColor("#0097a7"),
			M800: setColor("#00838f"),
			M900: setColor("#006064"),
			A100: setColor("#84ffff"),
			A200: setColor("#18ffff"),
			A400: setColor("#00e5ff"),
			A700: setColor("#00b8d4"),
		},
		Teal: &colorPaletteA{
			M50:  setColor("#e0f2f1"),
			M100: setColor("#b2dfdb"),
			M200: setColor("#80cbc4"),
			M300: setColor("#4db6ac"),
			M400: setColor("#26a69a"),
			M500: setColor("#009688"),
			M600: setColor("#00897b"),
			M700: setColor("#00796b"),
			M800: setColor("#00695c"),
			M900: setColor("#004d40"),
			A100: setColor("#a7ffeb"),
			A200: setColor("#64ffda"),
			A400: setColor("#1de9b6"),
			A700: setColor("#00bfa5"),
		},
		Green: &colorPaletteA{
			M50:  setColor("#e8f5e9"),
			M100: setColor("#c8e6c9"),
			M200: setColor("#a5d6a7"),
			M300: setColor("#81c784"),
			M400: setColor("#66bb6a"),
			M500: setColor("#4caf50"),
			M600: setColor("#43a047"),
			M700: setColor("#388e3c"),
			M800: setColor("#2e7d32"),
			M900: setColor("#1b5e20"),
			A100: setColor("#b9f6ca"),
			A200: setColor("#69f0ae"),
			A400: setColor("#00e676"),
			A700: setColor("#00c853"),
		},
		LightGreen: &colorPaletteA{
			M50:  setColor("#f1f8e9"),
			M100: setColor("#dcedc8"),
			M200: setColor("#c5e1a5"),
			M300: setColor("#aed581"),
			M400: setColor("#9ccc65"),
			M500: setColor("#8bc34a"),
			M600: setColor("#7cb342"),
			M700: setColor("#689f38"),
			M800: setColor("#558b2f"),
			M900: setColor("#33691e"),
			A100: setColor("#ccff90"),
			A200: setColor("#b2ff59"),
			A400: setColor("#76ff03"),
			A700: setColor("#64dd17"),
		},
		Lime: &colorPaletteA{
			M50:  setColor("#f9fbe7"),
			M100: setColor("#f0f4c3"),
			M200: setColor("#e6ee9c"),
			M300: setColor("#dce775"),
			M400: setColor("#d4e157"),
			M500: setColor("#cddc39"),
			M600: setColor("#c0ca33"),
			M700: setColor("#afb42b"),
			M800: setColor("#9e9d24"),
			M900: setColor("#827717"),
			A100: setColor("#f4ff81"),
			A200: setColor("#eeff41"),
			A400: setColor("#c6ff00"),
			A700: setColor("#aeea00"),
		},
		Yellow: &colorPaletteA{
			M50:  setColor("#fffde7"),
			M100: setColor("#fff9c4"),
			M200: setColor("#fff59d"),
			M300: setColor("#fff176"),
			M400: setColor("#ffee58"),
			M500: setColor("#ffeb3b"),
			M600: setColor("#fdd835"),
			M700: setColor("#fbc02d"),
			M800: setColor("#f9a825"),
			M900: setColor("#f57f17"),
			A100: setColor("#ffff8d"),
			A200: setColor("#ffff00"),
			A400: setColor("#ffea00"),
			A700: setColor("#ffd600"),
		},
		Amber: &colorPaletteA{
			M50:  setColor("#fff8e1"),
			M100: setColor("#ffecb3"),
			M200: setColor("#ffe082"),
			M300: setColor("#ffd54f"),
			M400: setColor("#ffca28"),
			M500: setColor("#ffc107"),
			M600: setColor("#ffb300"),
			M700: setColor("#ffa000"),
			M800: setColor("#ff8f00"),
			M900: setColor("#ff6f00"),
			A100: setColor("#ffe57f"),
			A200: setColor("#ffd740"),
			A400: setColor("#ffc400"),
			A700: setColor("#ffab00"),
		},
		Orange: &colorPaletteA{
			M50:  setColor("#fff3e0"),
			M100: setColor("#ffe0b2"),
			M200: setColor("#ffcc80"),
			M300: setColor("#ffb74d"),
			M400: setColor("#ffa726"),
			M500: setColor("#ff9800"),
			M600: setColor("#fb8c00"),
			M700: setColor("#f57c00"),
			M800: setColor("#ef6c00"),
			M900: setColor("#e65100"),
			A100: setColor("#ffd180"),
			A200: setColor("#ffab40"),
			A400: setColor("#ff9100"),
			A700: setColor("#ff6d00"),
		},
		DeepOrange: &colorPaletteA{
			M50:  setColor("#fbe9e7"),
			M100: setColor("#ffccbc"),
			M200: setColor("#ffab91"),
			M300: setColor("#ff8a65"),
			M400: setColor("#ff7043"),
			M500: setColor("#ff5722"),
			M600: setColor("#f4511e"),
			M700: setColor("#e64a19"),
			M800: setColor("#d84315"),
			M900: setColor("#bf360c"),
			A100: setColor("#ff9e80"),
			A200: setColor("#ff6e40"),
			A400: setColor("#ff3d00"),
			A700: setColor("#dd2c00"),
		},
	}
}
