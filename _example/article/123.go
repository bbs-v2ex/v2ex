package main

import (
	"fmt"
	"v2ex/until"
)

func main() {

	for {
		avatar := until.RandomAvatar()
		fmt.Println(avatar)
	}
}
