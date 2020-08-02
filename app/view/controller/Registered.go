package controller

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view"
)

func Registered(c *gin.Context) {
	_ht := defaultData(c)
	view.Render(c, "_registered", _ht)
}
