package v2ex

import (
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/foolin/goview/supports/gorice"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"v2ex/config"
	"v2ex/view_func"
)

func Run() {
	cg, err := config.LoadingConfigSourceFile()
	if err != nil {
		log.Fatal("加载配置文件失败", err)
		return
	}
	r := &gin.Engine{}

	debug := cg.Run.Debug

	//设置 gin 启动参数
	if debug {
		r = gin.Default()
	} else {
		r = gin.New()
	}

	//处理静态文件
	if debug {
		r.Static("/static", "./server/z_static")
	} else {
		box := rice.MustFindBox("z_static")
		cssFileServer := http.StripPrefix("/static", http.FileServer(box.HTTPBox()))
		r.GET("/static/:a", gin.WrapH(cssFileServer))
	}

	//加载全局模板函数
	tempFunc := view_func.TempFunc()

	//模板文件是否 打包
	if debug {
		_view_config := goview.DefaultConfig
		_view_config.Root = "server/z_view"
		_view_config.Funcs = tempFunc
		_view_config.DisableCache = true
		r.HTMLRender = ginview.New(_view_config)

	} else {
		basic := gorice.NewWithConfig(rice.MustFindBox("z_view"), goview.Config{
			Root:         "z_view",
			Extension:    ".html",
			Master:       "layouts/master",
			Funcs:        tempFunc,
			DisableCache: true,
		})
		r.HTMLRender = ginview.Wrap(basic)
	}
	//监听端口启动

	err = r.Run(fmt.Sprintf("%s:%d", cg.Run.LocaIP, cg.Run.Port))
	if err != nil {
		log.Fatal("启动失败", err)
		return
	}
}
