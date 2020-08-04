package router

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/model"
)

func checkLogin(c *gin.Context) {
	token := c.GetHeader("____token")
	//根据 Token 查询本人是否登录
	member_token := model.MemberToken{}
	mc.Table(member_token.Table()).Where(bson.M{"token": token}).FindOne(&member_token)
	if member_token.MID == 0 {
		result_json := c_code.V1GinError(500, "未登录")
		c.JSON(200, result_json)
		c.Abort()
		return
	}
	c.Set("mid", member_token.MID)
}
