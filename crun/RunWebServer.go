package crun

import (
	"fmt"
	"github.com/123456/c_code/mc"
	rice "github.com/GeertJohan/go.rice"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/foolin/goview/supports/gorice"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"v2ex/app/api/manage"
	api_router "v2ex/app/api/router"
	"v2ex/app/view"
	view_router "v2ex/app/view/router"
	"v2ex/config"
	"v2ex/model"
	"v2ex/until"
	"v2ex/view_func"
)

func RunWebServer() {
	var err error

	cg := config.GetConfig()
	manage.Init()
	r := &gin.Engine{}

	debug := cg.Run.Debug

	//设置 gin 启动参数
	if debug {
		r = gin.Default()
	} else {
		r = gin.New()
	}

	r.Static("/tmp", cg.Run.TempUploadDir)
	//处理静态文件
	if !config.BuildStaticAndTemplate {
		r.Static("/static", "./app/view/view_static")
	} else {
		box := rice.MustFindBox("../app/view/view_static")
		cssFileServer := http.StripPrefix("/static", http.FileServer(box.HTTPBox()))
		r.GET("/static/*a", gin.WrapH(cssFileServer))
	}

	//加载全局模板函数
	tempFunc := view_func.TempFunc()

	//模板文件是否 打包

	if !config.BuildStaticAndTemplate {
		_view_config := goview.DefaultConfig
		_view_config.Root = cg.ExecPath + "/app/view/view_template"
		_view_config.Funcs = tempFunc
		_view_config.DisableCache = true
		engine := ginview.New(_view_config)
		view.ViewEngine = engine
		r.HTMLRender = engine
	} else {
		basic := gorice.NewWithConfig(rice.MustFindBox("../app/view/view_template"), goview.Config{
			Root:         "view_template",
			Extension:    ".html",
			Master:       "layouts/master",
			Funcs:        tempFunc,
			DisableCache: true,
		})
		engine := ginview.Wrap(basic)
		view.ViewEngine = engine
		r.HTMLRender = engine
	}
	//监听端口启动

	//注册接口路由
	api_router.RegisterRoute(r)

	//注册普通路由一般用于页面展示
	view_router.RegisterRoute(r)

	//检测是否存在超级管理员 如果没有则直接创建
	root_member := model.Member{}
	mc.Table(root_member.Table()).Where(bson.M{"member_type": model.MemberTypeRoot}).FindOne(&root_member)

	if root_member.ID.Hex() == mc.Empty {
		member := model.Member{
			ID:         primitive.NewObjectID(),
			MID:        9999999999,
			UserName:   "root",
			MemberType: model.MemberTypeRoot,
			Avatar:     until.RandomAvatar(),
			IsUser:     false,
		}
		mc.Table(member.Table()).Insert(member)
		//写入内容表
		member_more := model.MemberMore{
			ID:       member.ID,
			PassWord: member.EncryptionPassWord("root"),
		}
		err = mc.Table(member_more.Table()).Insert(member_more)
	}

	err = r.Run(fmt.Sprintf("%s:%d", cg.Run.LocaIP, cg.Run.Port))
	if err != nil {
		log.Fatal("启动失败", err)
		return
	}
}
