package manage

import (
	"github.com/123456/c_code"
	"github.com/gin-gonic/gin"
	"v2ex/app/api"
)

func MemberNav(c *gin.Context) {
	user_info := api.GetNowUserInfo(c)
	if user_info.MemberType == 1 {
		u := "/_/config/"
		list := []gin.H{
			{
				"t": "SEO设置",
				"u": u + "seo",
			},
		}
		//登录成功
		result_json := c_code.V1GinSuccess(list)
		c.JSON(200, result_json)
		return
	} else {
		result_json := c_code.V1GinError(104, "鉴权？")
		c.JSON(200, result_json)
		return
	}
}
