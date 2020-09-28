package root_api

import (
	"github.com/123456/c_code"
	"github.com/gin-gonic/gin"
	"v2ex/app/nc"
	"v2ex/model"
)

func apiAuth(c *gin.Context) {
	result_json := c_code.V1GinSuccess(model.SiteConfig{}.GetApiAuth())
	c.JSON(200, result_json)
}
func apiAuthPost(c *gin.Context) {
	sc := model.SiteConfigApiAuth{}
	c.BindJSON(&sc)

	err := model.SiteConfig{}.SetApiAuth(sc)
	if err != nil {
		result_json := c_code.V1GinError(101, err.Error())
		c.JSON(200, result_json)
		return
	}
	nc.ReloadConfig()
	result_json := c_code.V1GinSuccess("", "修改成功")
	c.JSON(200, result_json)
}
