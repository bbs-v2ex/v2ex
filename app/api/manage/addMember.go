package manage

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/model"
)

type _add_member_post struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func AddMember(c *gin.Context) {
	_f := _add_member_post{}
	c.BindJSON(&_f)

	mid, err := model.AutoID{}.MemberID()
	if err != nil {
		result_json := c_code.V1GinError(101, "获取会员ID失败")
		c.JSON(200, result_json)
		return
	}

	member := model.Member{
		ID:       primitive.NewObjectID(),
		MID:      mid,
		UserName: _f.UserName,
		Avatar:   "",
	}
	err = mc.Table(member.Table()).Insert(member)
	if err != nil {
		result_json := c_code.V1GinError(102, "写入索引表失败")
		c.JSON(200, result_json)
		return
	}
	//写入内容表
	member_more := model.MemberMore{
		ID:       primitive.NewObjectID(),
		MID:      mid,
		PassWord: member.EncryptionPassWord(_f.PassWord),
	}
	err = mc.Table(member_more.Table()).Insert(member_more)
	if err != nil {
		result_json := c_code.V1GinError(103, "写入详细数据失败")
		c.JSON(200, result_json)
		return
	}
	result_json := c_code.V1GinSuccess("", "注册成功,前去登录", "/login")
	c.JSON(200, result_json)
}
