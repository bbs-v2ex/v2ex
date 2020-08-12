package main

import (
	"fmt"
	"path/filepath"
	"v2ex"
	"v2ex/app/nc"
	"v2ex/config"
)

func main() {
	abs, err := filepath.Abs("./")
	fmt.Println(abs, err)
	//链接数据库
	v2ex.ConnectMongodb()

	//加载SEO 配置
	nc.Init()

	fmt.Println(config.GetConfig().Dump())
	//启动 web 服务器
	v2ex.RunWebServer()
	//打印配置
}
