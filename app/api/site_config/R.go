package site_config

import "github.com/gin-gonic/gin"

func R(r *gin.RouterGroup) {
	r1 := r.Group("/config")
	r1.GET("/seo", seo)
	r1.POST("/seo", seopost)
}
