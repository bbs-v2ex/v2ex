package member

import (
	"github.com/123456/c_code"
	"github.com/gin-gonic/gin"
	"v2ex/app/api"
	"v2ex/model"
)

func getUserInfo(c *gin.Context) {
	mid := api.GetMID(c)

	user_info := model.Member{}.GetUserInfo(mid, true)

	result := _set_user_info{
		UserName:    user_info.UserName,
		Avatar:      model.Avatar(user_info.Avatar),
		Des:         user_info.More.Des,
		DesDetailed: user_info.More.DesDetailed,
	}
	result_json := c_code.V1GinSuccess(result)
	c.JSON(200, result_json)
}
