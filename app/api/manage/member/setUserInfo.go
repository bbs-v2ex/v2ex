package member

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/app/api"
	"v2ex/model"
)

type _set_user_info struct {
	UserName    string `json:"user_name" bson:"user_name"`
	Avatar      string `json:"avatar" bson:"avatar"`
	Des         string `json:"des" bson:"des"`
	DesDetailed string `json:"des_detailed" bson:"des_detailed"`
}

func setUserInfo(c *gin.Context) {
	_f := _set_user_info{}
	c.BindJSON(&_f)

	//查询自己得ID
	mid := api.GetMID(c)
	member := model.Member{}
	mc.Table(member.Table()).Where(bson.M{"mid": mid}).FindOne(&member)
	if member.ID.Hex() == mc.Empty {
		result_json := c_code.V1GinError(101, "修改失败")
		c.JSON(200, result_json)
		return
	}
	put_member_more := bson.M{
		"des":          _f.Des,
		"des_detailed": _f.DesDetailed,
	}
	err := mc.Table(model.MemberMore{}.Table()).Where(bson.M{"_id": member.ID}).UpdateOne(put_member_more)
	if err != nil {
		result_json := c_code.V1GinError(102, "修改失败")
		c.JSON(200, result_json)
		return
	}
	result_json := c_code.V1GinSuccess("", "修改成功")
	c.JSON(200, result_json)
	return
}
