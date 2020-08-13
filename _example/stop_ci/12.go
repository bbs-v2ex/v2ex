package main

import (
	"fmt"
	"github.com/123456/c_code"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"v2ex/until"
)

func main() {

	chiReg := regexp.MustCompile("[^\u4e00-\u9fa5]")
	log.Println(chiReg.ReplaceAllString("haha 哈哈东方科技dfjkdj2323", ""))

	ci_list := until.StopCIFilter("你身边的学霸都有怎样的学习方法或习惯")
	fmt.Println(ci_list)
	//list := StopCI()
	//r := regexp.MustCompile(list).ReplaceAllString(`你身边的学霸都有怎样的学习方法或习惯？`, ",")
	//fmt.Println(r)
}

func StopCI() string {
	stop_list := []string{}
	path := `C:\Users\Administrator\Desktop\停用词\`
	dir, _ := ioutil.ReadDir(path)
	for _, file := range dir {
		if strings.HasSuffix(file.Name(), ".txt") {
			readFile, err := ioutil.ReadFile(path + file.Name())
			if err != nil {
				continue
			}
			s1 := string(readFile)
			s := strings.ReplaceAll(string(s1), "\r", "")
			list := strings.Split(s, "\n")
			stop_list = append(stop_list, list...)
			fmt.Println(len(stop_list))
		}
	}
	stop_list = c_code.UniqueSliceString(stop_list)
	must := "[" + strings.Join(stop_list, "|") + "]"
	must = strings.ReplaceAll(must, "--", "\\-\\-")
	must = strings.ReplaceAll(must, "+", "\\+")
	//fmt.Println(must)
	return must
}
