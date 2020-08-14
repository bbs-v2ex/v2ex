package site_config

import (
	"github.com/123456/c_code"
	"github.com/gin-gonic/gin"
	"v2ex/app/api"
)

func R(r *gin.RouterGroup) {
	r1 := r.Group("/config")
	r1.Use(isok)
	r1.GET("/seo", seo)
	r1.POST("/seo", seopost)

	r1.GET("/api_auth", api_auth)
	r1.POST("/api_auth", api_auth_post)
	r1.GET("/db_index", db_index)
	r1.GET("/create_index", create_index)
}

func isok(c *gin.Context) {
	user := api.GetNowUserInfo(c)
	if user.MemberType != 1 {
		result := c_code.V1GinError(500, "没权限啊")
		c.JSON(200, result)
		c.Abort()
		return
	}
}
