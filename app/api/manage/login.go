package manage

import (
	"fmt"
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/app/nc"
	"v2ex/model"
	"v2ex/until"
)

func Login(c *gin.Context) {
	_f := _add_member_post{}
	c.BindJSON(&_f)
	api_auth := nc.GetApiAuth()
	member := model.Member{}
	if api_auth.SpiderSign == _f.Sign && api_auth.SpiderSign != "" {
		list_mid := []model.Member{}
		if _f.MID != 0 {
			mc.Table(model.Member{}.Table()).Where(bson.M{"is_user": false, "mid": _f.MID}).Projection(bson.M{"mid": 1}).FindOne(&member)
		}
		if member.MID == 0 {
			mc.Table(model.Member{}.Table()).Where(bson.M{"is_user": false}).Projection(bson.M{"mid": 1}).Find(&list_mid)
			member = list_mid[c_code.Rand(len(list_mid)-1)]
		}

	} else {
		if _f.UserName == "" || _f.PassWord == "" {
			result_json := c_code.V1GinError(101, "用户或密码错误")
			c.JSON(200, result_json)
			return
		}
		//查找会员名

		mc.Table(member.Table()).Where(bson.M{"user_name": _f.UserName}).FindOne(&member)
		if member.UserName == "" {
			result_json := c_code.V1GinError(102, "用户或密码错误")
			c.JSON(200, result_json)
			return
		}
		//查询密码
		member_more := model.MemberMore{}
		mc.Table(member_more.Table()).Where(bson.M{"_id": member.ID}).FindOne(&member_more)
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
	}

	//生成uuid 唯一值
	v4, err := uuid.NewV4()
	if err != nil {
		result_json := c_code.V1GinError(105, "生成token失败")
		c.JSON(200, result_json)
		return
	}

	token := v4.String()

	//讲本次登录信息存入数据库
	last_time := until.MemberTokenAddValidPeriod()
	fmt.Println(last_time)
	login_token := model.MemberToken{
		ID:     primitive.NewObjectID(),
		MID:    member.MID,
		Token:  token,
		Expire: last_time,
	}
	count, err := mc.Table(login_token.Table()).Where(bson.M{"mid": member.MID}).Count()
	if count > 10 {
		result_json := c_code.V1GinError(106, "已存在 10 个会话,已限制进行登录")
		c.JSON(200, result_json)
		return
	}

	mc.Table(login_token.Table()).Insert(&login_token)

	//登录成功
	result_json := c_code.V1GinSuccess("", "登录成功")
	result_json["token"] = token

	c.JSON(200, result_json)
}
