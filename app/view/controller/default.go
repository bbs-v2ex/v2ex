package controller

import (
	"github.com/gin-gonic/gin"
	"v2ex/config"
	"v2ex/model"
)

func DefaultData(c *gin.Context) (_ht gin.H) {
	_con := config.GetConfig()
	_ht = gin.H{}
	_ht["_______API"] = "/api/manage"
	_ht["___upload_server"] = _con.Run.UploadServer
	//初始化tdk
	seo := model.SiteConfig{}.GetSeo()
	_ht["t"] = seo.T
	_ht["d"] = seo.D
	_ht["k"] = seo.K
	_ht["t_"] = seo.T_
	return
}

func defaultData(c *gin.Context) gin.H {
	return DefaultData(c)
}
