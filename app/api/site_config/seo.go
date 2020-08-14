package site_config

import (
	"github.com/123456/c_code"
	"github.com/gin-gonic/gin"
	"v2ex/app/nc"
	"v2ex/model"
)

func seo(c *gin.Context) {
	result_json := c_code.V1GinSuccess(model.SiteConfig{}.GetSeo())
	c.JSON(200, result_json)
}
func seopost(c *gin.Context) {
	sc := model.SiteConfigSeo{}
	c.BindJSON(&sc)

	err := model.SiteConfig{}.SetSeo(sc)
	if err != nil {
		result_json := c_code.V1GinError(101, err.Error())
		c.JSON(200, result_json)
		return
	}
	nc.ReloadSeoConfig()
	result_json := c_code.V1GinSuccess("", "修改成功")
	c.JSON(200, result_json)
}
