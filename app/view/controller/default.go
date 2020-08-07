package controller

import (
	"github.com/gin-gonic/gin"
	"v2ex/config"
)

func DefaultData(c *gin.Context) (_ht gin.H) {
	_con := config.GetConfig()
	_ht = gin.H{}
	_ht["_______API"] = "/api/manage"
	_ht["___upload_server"] = _con.Run.UploadServer
	return
}

func defaultData(c *gin.Context) gin.H {
	return DefaultData(c)
}
