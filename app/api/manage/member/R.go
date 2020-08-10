package member

import (
	"github.com/123456/c_code"
	"github.com/gin-gonic/gin"
	"v2ex/app/api"
)

func R(r *gin.RouterGroup) {
	r1 := r.Group("/member")
	r1.POST("/get_user_info", get_user_info)
	r1.POST("/set_user_info", set_user_info)
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
