package acorntypes

type ColParams struct {
	Content string
	Width   *string
}

type ButtonParams struct {
	Text, Link, HexColor, Align string
}

func String(s string) *string {
	return &s
}
