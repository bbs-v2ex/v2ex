package main

import (
	"fmt"
	"v2ex/config"
)

func main() {
	config.CreateConfigFile()
	file, err := config.LoadingConfigSourceFile()
	fmt.Println(err, file)
}
