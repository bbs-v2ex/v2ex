package main

import (
	"fmt"
	"regexp"
)

func main() {
	ref := "http://127.0.0.1:8777/a/15/r/5f2f2cd8c25bb5967b84ae6c"
	fmt.Println(regexp.MustCompile(`/r/[\w|\s]{24}`).ReplaceAllString(ref, ""))
}
