package acornmail

import (
	"fmt"

	"github.com/arskang/gomail-acorn-template/acorn"
)

type acornEmail struct {
	HTML *acorn.HTML
}

func NewAcornEmail() *acornEmail {
	return &acornEmail{
		HTML: &acorn.HTML{},
	}
}

func Example() {
	acorn := NewAcornEmail()
	row := acorn.HTML.GetRow("Hola")
	body := acorn.HTML.GetBoilerPlate(row)
	fmt.Println(body)
}
