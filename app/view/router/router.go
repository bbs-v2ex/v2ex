package router

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view/controller"
)

func RegisterRoute(r *gin.Engine) {
	r.GET("/", controller.Home)

	//注册
	r.GET("/registered", controller.Registered)

	//登录
	r.GET("/login", controller.Login)
}
