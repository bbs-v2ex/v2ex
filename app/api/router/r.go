package router

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/api/manage"
)

func RegisterRoute(r *gin.Engine) {
	ro := r.Group("/api/manage")
	ro.POST("/add_member", manage.AddMember)
}
