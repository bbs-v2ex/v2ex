package member

import (
	"github.com/123456/c_code"
	"github.com/gin-gonic/gin"
	"v2ex/app/api"
	"v2ex/app/common"
)

func get_user_info(c *gin.Context) {
	user_info := api.GetNowUserInfo(c)
	result := gin.H{
		"user_name":    user_info.UserName,
		"avatar":       common.Avatar(user_info.Avatar),
		"des":          user_info.More.Des,
		"des_detailed": user_info.More.DesDetailed,
	}
	result_json := c_code.V1GinSuccess(result)
	c.JSON(200, result_json)
}
