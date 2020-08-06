package controller

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view"
)

func Home(c *gin.Context) {
	_ht := defaultData(c)
	view.Render(c, "index", _ht)
}
