package router

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view/controller"
)

func RegisterRoute(r *gin.Engine) {
	r.GET("/", controller.Home)

	r.GET("/registered", controller.Registered)
}
