package root_api

import (
	"github.com/123456/c_code"
	"github.com/gin-gonic/gin"
	"v2ex/app/api"
	"v2ex/model"
)

func R(r *gin.RouterGroup) {
	r1 := r.Group("/config")
	r1.Use(isRootAuth)
	r1.GET("/seo", seo)
	r1.POST("/seo", seoPost)
	r1.GET("/api_auth", apiAuth)
	r1.POST("/api_auth", apiAuthPost)
	r1.GET("/db_index", dbIndex)
	r1.GET("/create_index", createIndex)
	//数据审核结果
	r1.POST("/data_check", dataCheck)

	//数据管理页面
	r1.POST("/data_search_nav", dataSearchNav)
	r1.POST("/data_search", dataSearch)
	r1.POST("/data_del", dataDel)

}

func isRootAuth(c *gin.Context) {
	user := api.GetNowUserInfo(c)
	if user.MemberType != model.MemberTypeRoot {
		result := c_code.V1GinError(5000, "没权限啊")
		c.JSON(200, result)
		c.Abort()
		return
	}
}
