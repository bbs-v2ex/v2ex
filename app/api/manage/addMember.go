package manage

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"v2ex/app/api"
	"v2ex/app/nc"
	"v2ex/model"
	"v2ex/until"
)

type _add_member_post struct {
	UserName string        `json:"username" validate:"required" common:"用户名"`
	PassWord string        `json:"password" validate:"required" common:"用户名"`
	Sign     string        `json:"sign"`
	MID      model.MIDTYPE `json:"mid"`
}

func usernameisok(username string) bool {
	member := model.Member{}
	mc.Table(member.Table()).Where(bson.M{"user_name": username}).FindOne(&member)
	if member.ID.Hex() == mc.Empty {
		return true
	}
	return false
}

func AddMember(c *gin.Context) {
	api_auth := nc.GetApiAuth()
	if !api_auth.Register {
		result_json := c_code.V1GinError(100, "网站关闭注册")
		c.JSON(200, result_json)
		return
	}
	_f := _add_member_post{}
	c.BindJSON(&_f)

	validator := api.VerifyValidator(_f)

	if validator != "" {
		result_json := c_code.V1GinError(101, validator)
		c.JSON(200, result_json)
		return
	}

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
		Avatar:   until.RandomAvatar(),
		IsUser:   true,
	}

	//查看名字是否ok
	if !usernameisok(_f.UserName) {
		result_json := c_code.V1GinError(101, "用户名出现问题，请更换一个")
		c.JSON(200, result_json)
		return
	}

	err = mc.Table(member.Table()).Insert(member)
	if err != nil {
		result_json := c_code.V1GinError(102, "写入索引表失败")
		c.JSON(200, result_json)
		return
	}
	//写入内容表
	member_more := model.MemberMore{
		ID:           member.ID,
		PassWord:     member.EncryptionPassWord(_f.PassWord),
		RegisterTime: time.Now(),
	}
	err = mc.Table(member_more.Table()).Insert(member_more)
	if err != nil {
		result_json := c_code.V1GinError(103, "写入详细数据失败")
		c.JSON(200, result_json)
		return
	}
	result_json := c_code.V1GinSuccess("", "注册成功,前去登录", "/login")
	model.AutoID{}.MemberAdd()
	c.JSON(200, result_json)
}
