package acornstyles

import (
	"fmt"
	"regexp"

	"github.com/arskang/gomail-acorn-template/acorntypes"
)

type colorPalette struct {
	_50, _100, _200, _300, _400, _500, _600, _700, _800, _900 *acorntypes.Color
}

type colorPaletteA struct {
	_50, _100, _200, _300, _400, _500, _600, _700, _800, _900, A100, A200, A400, A700 *acorntypes.Color
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

func NewAcornColor(hex string) (*acorntypes.Color, error) {
	match, _ := regexp.MatchString(RGX_HEXCOLOR, hex)
	if match {
		return setColor(hex), nil
	}
	return nil, fmt.Errorf("it is not a hexadecimal color")
}

func GetColors() *colors {
	return &colors{
		White: setColor("#ffffff"),
		Black: setColor("#21292f"),
		Brown: &colorPalette{
			_50:  setColor("#efebe9"),
			_100: setColor("#d7ccc8"),
			_200: setColor("#bcaaa4"),
			_300: setColor("#a1887f"),
			_400: setColor("#8d6e63"),
			_500: setColor("#795548"),
			_600: setColor("#6d4c41"),
			_700: setColor("#5d4037"),
			_800: setColor("#4e342e"),
			_900: setColor("#3e2723"),
		},
		Grey: &colorPalette{
			_50:  setColor("#fafafa"),
			_100: setColor("#d7ccc8"),
			_200: setColor("#eeeeee"),
			_300: setColor("#e0e0e0"),
			_400: setColor("#bdbdbd"),
			_500: setColor("#9e9e9e"),
			_600: setColor("#757575"),
			_700: setColor("#616161"),
			_800: setColor("#424242"),
			_900: setColor("#212121"),
		},
		BlueGrey: &colorPalette{
			_50:  setColor("#eceff1"),
			_100: setColor("#cfd8dc"),
			_200: setColor("#b0bec5"),
			_300: setColor("#90a4ae"),
			_400: setColor("#78909c"),
			_500: setColor("#607d8b"),
			_600: setColor("#546e7a"),
			_700: setColor("#455a64"),
			_800: setColor("#37474f"),
			_900: setColor("#263238"),
		},
		Red: &colorPaletteA{
			_50:  setColor("#ffebee"),
			_100: setColor("#ffcdd2"),
			_200: setColor("#ef9a9a"),
			_300: setColor("#e57373"),
			_400: setColor("#ef5350"),
			_500: setColor("#f44336"),
			_600: setColor("#e53935"),
			_700: setColor("#d32f2f"),
			_800: setColor("#c62828"),
			_900: setColor("#b71c1c"),
			A100: setColor("#ff5252"),
			A200: setColor("#ff5252"),
			A400: setColor("#ff1744"),
			A700: setColor("#d50000"),
		},
		Pink: &colorPaletteA{
			_50:  setColor("#fce4ec"),
			_100: setColor("#f8bbd0"),
			_200: setColor("#f48fb1"),
			_300: setColor("#f06292"),
			_400: setColor("#ec407a"),
			_500: setColor("#e91e63"),
			_600: setColor("#d81b60"),
			_700: setColor("#c2185b"),
			_800: setColor("#ad1457"),
			_900: setColor("#880e4f"),
			A100: setColor("#ff80ab"),
			A200: setColor("#ff4081"),
			A400: setColor("#f50057"),
			A700: setColor("#c51162"),
		},
		Purple: &colorPaletteA{
			_50:  setColor("#f3e5f5"),
			_100: setColor("#e1bee7"),
			_200: setColor("#ce93d8"),
			_300: setColor("#ba68c8"),
			_400: setColor("#ab47bc"),
			_500: setColor("#9c27b0"),
			_600: setColor("#8e24aa"),
			_700: setColor("#7b1fa2"),
			_800: setColor("#6a1b9a"),
			_900: setColor("#4a148c"),
			A100: setColor("#ea80fc"),
			A200: setColor("#e040fb"),
			A400: setColor("#d500f9"),
			A700: setColor("#aa00ff"),
		},
		DeepPurple: &colorPaletteA{
			_50:  setColor("#ede7f6"),
			_100: setColor("#d1c4e9"),
			_200: setColor("#b39ddb"),
			_300: setColor("#9575cd"),
			_400: setColor("#7e57c2"),
			_500: setColor("#673ab7"),
			_600: setColor("#5e35b1"),
			_700: setColor("#512da8"),
			_800: setColor("#4527a0"),
			_900: setColor("#311b92"),
			A100: setColor("#b388ff"),
			A200: setColor("#7c4dff"),
			A400: setColor("#651fff"),
			A700: setColor("#6200ea"),
		},
		Indigo: &colorPaletteA{
			_50:  setColor("#e8eaf6"),
			_100: setColor("#c5cae9"),
			_200: setColor("#9fa8da"),
			_300: setColor("#7986cb"),
			_400: setColor("#5c6bc0"),
			_500: setColor("#3f51b5"),
			_600: setColor("#3949ab"),
			_700: setColor("#303f9f"),
			_800: setColor("#283593"),
			_900: setColor("#1a237e"),
			A100: setColor("#8c9eff"),
			A200: setColor("#536dfe"),
			A400: setColor("#3d5afe"),
			A700: setColor("#304ffe"),
		},
		Blue: &colorPaletteA{
			_50:  setColor("#e3f2fd"),
			_100: setColor("#bbdefb"),
			_200: setColor("#90caf9"),
			_300: setColor("#64b5f6"),
			_400: setColor("#42a5f5"),
			_500: setColor("#2196f3"),
			_600: setColor("#1e88e5"),
			_700: setColor("#1976d2"),
			_800: setColor("#1565c0"),
			_900: setColor("#0d47a1"),
			A100: setColor("#82b1ff"),
			A200: setColor("#448aff"),
			A400: setColor("#2979ff"),
			A700: setColor("#2962ff"),
		},
		LightBlue: &colorPaletteA{
			_50:  setColor("#e1f5fe"),
			_100: setColor("#b3e5fc"),
			_200: setColor("#81d4fa"),
			_300: setColor("#4fc3f7"),
			_400: setColor("#29b6f6"),
			_500: setColor("#03a9f4"),
			_600: setColor("#039be5"),
			_700: setColor("#0288d1"),
			_800: setColor("#0277bd"),
			_900: setColor("#01579b"),
			A100: setColor("#80d8ff"),
			A200: setColor("#40c4ff"),
			A400: setColor("#00b0ff"),
			A700: setColor("#0091ea"),
		},
		Cyan: &colorPaletteA{
			_50:  setColor("#e0f7fa"),
			_100: setColor("#b2ebf2"),
			_200: setColor("#80deea"),
			_300: setColor("#4dd0e1"),
			_400: setColor("#26c6da"),
			_500: setColor("#00bcd4"),
			_600: setColor("#00acc1"),
			_700: setColor("#0097a7"),
			_800: setColor("#00838f"),
			_900: setColor("#006064"),
			A100: setColor("#84ffff"),
			A200: setColor("#18ffff"),
			A400: setColor("#00e5ff"),
			A700: setColor("#00b8d4"),
		},
		Teal: &colorPaletteA{
			_50:  setColor("#e0f2f1"),
			_100: setColor("#b2dfdb"),
			_200: setColor("#80cbc4"),
			_300: setColor("#4db6ac"),
			_400: setColor("#26a69a"),
			_500: setColor("#009688"),
			_600: setColor("#00897b"),
			_700: setColor("#00796b"),
			_800: setColor("#00695c"),
			_900: setColor("#004d40"),
			A100: setColor("#a7ffeb"),
			A200: setColor("#64ffda"),
			A400: setColor("#1de9b6"),
			A700: setColor("#00bfa5"),
		},
		Green: &colorPaletteA{
			_50:  setColor("#e8f5e9"),
			_100: setColor("#c8e6c9"),
			_200: setColor("#a5d6a7"),
			_300: setColor("#81c784"),
			_400: setColor("#66bb6a"),
			_500: setColor("#4caf50"),
			_600: setColor("#43a047"),
			_700: setColor("#388e3c"),
			_800: setColor("#2e7d32"),
			_900: setColor("#1b5e20"),
			A100: setColor("#b9f6ca"),
			A200: setColor("#69f0ae"),
			A400: setColor("#00e676"),
			A700: setColor("#00c853"),
		},
		LightGreen: &colorPaletteA{
			_50:  setColor("#f1f8e9"),
			_100: setColor("#dcedc8"),
			_200: setColor("#c5e1a5"),
			_300: setColor("#aed581"),
			_400: setColor("#9ccc65"),
			_500: setColor("#8bc34a"),
			_600: setColor("#7cb342"),
			_700: setColor("#689f38"),
			_800: setColor("#558b2f"),
			_900: setColor("#33691e"),
			A100: setColor("#ccff90"),
			A200: setColor("#b2ff59"),
			A400: setColor("#76ff03"),
			A700: setColor("#64dd17"),
		},
		Lime: &colorPaletteA{
			_50:  setColor("#f9fbe7"),
			_100: setColor("#f0f4c3"),
			_200: setColor("#e6ee9c"),
			_300: setColor("#dce775"),
			_400: setColor("#d4e157"),
			_500: setColor("#cddc39"),
			_600: setColor("#c0ca33"),
			_700: setColor("#afb42b"),
			_800: setColor("#9e9d24"),
			_900: setColor("#827717"),
			A100: setColor("#f4ff81"),
			A200: setColor("#eeff41"),
			A400: setColor("#c6ff00"),
			A700: setColor("#aeea00"),
		},
		Yellow: &colorPaletteA{
			_50:  setColor("#fffde7"),
			_100: setColor("#fff9c4"),
			_200: setColor("#fff59d"),
			_300: setColor("#fff176"),
			_400: setColor("#ffee58"),
			_500: setColor("#ffeb3b"),
			_600: setColor("#fdd835"),
			_700: setColor("#fbc02d"),
			_800: setColor("#f9a825"),
			_900: setColor("#f57f17"),
			A100: setColor("#ffff8d"),
			A200: setColor("#ffff00"),
			A400: setColor("#ffea00"),
			A700: setColor("#ffd600"),
		},
		Amber: &colorPaletteA{
			_50:  setColor("#fff8e1"),
			_100: setColor("#ffecb3"),
			_200: setColor("#ffe082"),
			_300: setColor("#ffd54f"),
			_400: setColor("#ffca28"),
			_500: setColor("#ffc107"),
			_600: setColor("#ffb300"),
			_700: setColor("#ffa000"),
			_800: setColor("#ff8f00"),
			_900: setColor("#ff6f00"),
			A100: setColor("#ffe57f"),
			A200: setColor("#ffd740"),
			A400: setColor("#ffc400"),
			A700: setColor("#ffab00"),
		},
		Orange: &colorPaletteA{
			_50:  setColor("#fff3e0"),
			_100: setColor("#ffe0b2"),
			_200: setColor("#ffcc80"),
			_300: setColor("#ffb74d"),
			_400: setColor("#ffa726"),
			_500: setColor("#ff9800"),
			_600: setColor("#fb8c00"),
			_700: setColor("#f57c00"),
			_800: setColor("#ef6c00"),
			_900: setColor("#e65100"),
			A100: setColor("#ffd180"),
			A200: setColor("#ffab40"),
			A400: setColor("#ff9100"),
			A700: setColor("#ff6d00"),
		},
		DeepOrange: &colorPaletteA{
			_50:  setColor("#fbe9e7"),
			_100: setColor("#ffccbc"),
			_200: setColor("#ffab91"),
			_300: setColor("#ff8a65"),
			_400: setColor("#ff7043"),
			_500: setColor("#ff5722"),
			_600: setColor("#f4511e"),
			_700: setColor("#e64a19"),
			_800: setColor("#d84315"),
			_900: setColor("#bf360c"),
			A100: setColor("#ff9e80"),
			A200: setColor("#ff6e40"),
			A400: setColor("#ff3d00"),
			A700: setColor("#dd2c00"),
		},
	}
}
