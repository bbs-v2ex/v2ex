package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
)

type T struct {
	Title   string `json:"title" valid:"required"`
	Content string `json:"content" `
	Text    string `json:"text" valid:"required"`
}

func main() {
	_f := T{
		Title:   "",
		Content: "",
		Text:    "",
	}

	_, err := govalidator.ValidateStruct(_f)
	fmt.Println(err)
	// IS nil
}
