package controller

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view"
)

func Login(c *gin.Context) {
	_ht := defaultData(c)
	view.Render(c, "_login", _ht)
}
