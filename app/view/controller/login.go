package controller

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view"
	"v2ex/model"
)

func Login(c *gin.Context) {
	_ht := defaultData(c)
	_ht["login_success"] = model.UrlViewMemberConfig
	view.Render(c, "_login", _ht)
}
