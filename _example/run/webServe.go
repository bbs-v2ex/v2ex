package main

import (
	"fmt"
	"path/filepath"
	"v2ex"
	"v2ex/config"
)

func main() {
	abs, err := filepath.Abs("./")
	fmt.Println(abs, err)
	//链接数据库
	v2ex.ConnectMongodb()

	fmt.Println(config.GetConfig().Dump())
	//启动 web 服务器
	v2ex.RunWebServer()
	//打印配置

}
