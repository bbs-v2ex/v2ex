package manage

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/model"
)

func Login(c *gin.Context) {
	_f := _add_member_post{}
	c.BindJSON(&_f)

	if _f.UserName == "" || _f.PassWord == "" {
		result_json := c_code.V1GinError(101, "用户或密码错误")
		c.JSON(200, result_json)
		return
	}
	//查找会员名
	member := model.Member{}
	mc.Table(member.Table()).Where(bson.M{"user_name": _f.UserName}).FindOne(&member)
	if member.UserName == "" {
		result_json := c_code.V1GinError(102, "用户或密码错误")
		c.JSON(200, result_json)
		return
	}
	//查询密码
	member_more := model.MemberMore{}
	mc.Table(member_more.Table()).Where(bson.M{"mid": member.MID}).FindOne(&member_more)
	if member_more.PassWord == "" {
		result_json := c_code.V1GinError(103, "用户或密码错误")
		c.JSON(200, result_json)
		return
	}
	if member.EncryptionPassWord(_f.PassWord) != member_more.PassWord {
		result_json := c_code.V1GinError(104, "用户或密码错误")
		c.JSON(200, result_json)
		return
	}
	//登录成功
	result_json := c_code.V1GinSuccess("", "登录成功")
	c.JSON(200, result_json)
}
