package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"v2ex/app/nc"
	"v2ex/config"
	"v2ex/model"
)

func DefaultData(c *gin.Context) (_ht gin.H) {
	seoconfig := nc.GetSeoConfig()
	_con := config.GetConfig()
	_ht = gin.H{}
	//页面类型
	_ht["_con"] = _con
	_ht["_______API"] = "/api/manage"
	_ht["___upload_server"] = _con.Run.UploadServer
	//初始化tdk
	_ht["_debug"] = _con.Run.Debug
	_ht["t"] = seoconfig.T
	_ht["d"] = seoconfig.D
	_ht["k"] = seoconfig.K
	_ht["t_"] = seoconfig.T_
	_ht["title_fgf"] = seoconfig.TitleDelimiter

	_ht["avatar"] = func(s string) string {
		return _con.Run.UploadServer + s
	}
	_ht["mu"] = model.UrlMember
	_ht["au"] = model.UrlArticle
	_ht["member_centre"] = model.UrlViewMemberConfig
	_ht["mumid"] = func(mid model.MIDTYPE) string {
		return fmt.Sprintf("/%s/%d", model.UrlTagMember, mid)
	}

	_ht["u_list"] = gin.H{
		"member":   "/" + model.UrlTagMember + "/",
		"article":  "/" + model.UrlTagArticle + "/",
		"question": "/" + model.UrlTagQuestion + "/",
	}

	//首页菜单
	_ht["home_nav"] = []gin.H{
		{
			"t": "动态",
			"u": "/last_activity",
		},
		{
			"t": "问题",
			"u": "/q/",
		},
		{
			"t": "文章",
			"u": "/a/",
		},
	}
	//公示信息
	_ht["seoconfig"] = seoconfig
	return
}

func defaultData(c *gin.Context) gin.H {
	return DefaultData(c)
}
