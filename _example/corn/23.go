package main

import (
	"fmt"
	"github.com/123456/c_code"
	"github.com/robfig/cron/v3"
)

func main() {
	//c := cron.New(cron.WithSeconds())
	c := cron.New()
	// 通过AddFunc注册

	_, err := c.AddFunc("*/1 * * * *", func() {
		//更新网站地图
		get, b, err := c_code.CGet("http://127.0.0.1:8777" + "/update_site_map")
		fmt.Println(get, b, err)
		//删除缓存的图片
	})
	fmt.Println(err)
	c.Start()
	select {}
}
