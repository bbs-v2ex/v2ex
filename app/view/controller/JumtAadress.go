package controller

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view"
)

func JumtAddress(c *gin.Context) {
	_ht := defaultData(c)
	_ht["_u"] = c.Query("u")
	view.Render(c, "_jumt_address", _ht)
}
