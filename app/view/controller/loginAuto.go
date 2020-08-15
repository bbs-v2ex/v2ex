package controller

import (
	"github.com/gin-gonic/gin"
)

func LoginAuto(c *gin.Context) {
	_ht := gin.H{}
	_ht["sign"] = c.Query("sign")
	_ht["mid"] = c.Query("mid")
	c.HTML(200, "_login_auto.html", _ht)
}
