package router

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/api/manage"
)

func RegisterRoute(r *gin.Engine) {

	//authMiddleware := init_jwt()

	ro := r.Group("/api/manage")
	ro.POST("/add_member", manage.AddMember)
	ro.POST("/login", manage.Login)

	ro.Use(checkLogin)
	ro.POST("/get_user_info", manage.GetUserInfo)
}
