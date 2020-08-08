package api

import (
	"github.com/123456/c_code"
	"strings"
)

func FilterTitle(title string) (_title string) {
	_title = c_code.RemoveHtmlTag(title)
	_title = strings.Trim(_title, "")
	return
}
func FilterContent(title string) (_title string) {
	_title = strings.Trim(title, "")
	return
}
