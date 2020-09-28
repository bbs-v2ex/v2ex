package main

import (
	"fmt"
	"log"
	"path/filepath"
	"v2ex/app/nc"
	"v2ex/config"
	"v2ex/crun"
)

func main() {
	abs, err := filepath.Abs("./")
	fmt.Println(abs, err)

	_, err = config.LoadingConfigSourceFile()
	if err != nil {
		log.Fatal("加载配置文件失败", err)
		return
	}

	config.BuildStaticAndTemplate = false
	/*build
	config.BuildStaticAndTemplate = true
	build*/

	//config.BuildStaticAndTemplate = true

	//链接数据库
	crun.ConnectMongodb()

	//加载SEO 配置
	nc.Init()

	fmt.Println(config.GetConfig().Dump())

	//系统定时任务
	crun.CornRun()

	//启动 web 服务器
	crun.RunWebServer()
	//打印配置
}
