package build

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TestBuild(t *testing.T) {
	file, err := ioutil.ReadFile("../../webServe.go")
	if err != nil {
		log.Println("编译失败", err.Error())
		return
	}
	f := string(file)
	f = strings.ReplaceAll(f, "/*build", "")
	f = strings.ReplaceAll(f, "build*/", "")
	ioutil.WriteFile("../../build/run.go", []byte(f), 0777)
}
