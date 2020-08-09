package article

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/app/api"
	"v2ex/app/common"
	"v2ex/model"
)

type _comment_child_list struct {
	RID primitive.ObjectID `json:"rid" validate:"len=12" comment:"评论ID"`
}

func comment_child_list(c *gin.Context) {
	_f := _comment_child_list{}
	c.BindJSON(&_f)
	validator := api.VerifyValidator(_f)
	if validator != "" {
		result_json := c_code.V1GinError(101, validator)
		c.JSON(200, result_json)
		return
	}
	//查询结果并返回
	_comment_child := []model.CommentChild{}
	mc.Table(model.CommentChild{}.Table()).Where(bson.M{"rid": _f.RID}).Limit(10).Find(&_comment_child)
	comment_child := []gin.H{}
	for _, v := range _comment_child {
		//获取text
		comment_text := model.CommentText{}
		mc.Table(comment_text.Table()).Where(bson.M{"_id": v.ID}).FindOne(&comment_text)
		if comment_text.ID.Hex() == mc.Empty {
			continue
		}
		//获取本身的信息
		find_user_chlient := model.Member{}
		self_member := find_user_chlient.GetUserInfo(v.MID, false)
		if self_member.MID == 0 {
			continue
		}
		child_one := gin.H{
			"u1": gin.H{
				"name":   self_member.UserName,
				"avatar": common.Avatar(self_member.Avatar),
				"mid":    self_member.MID,
			},
			"txt":  comment_text.Text,
			"time": c_code.StrTime(comment_text.ReleaseTime),
			"_id":  comment_text.ID.Hex(),
		}

		//检测是否实对别人进行回复
		if v.PID.Hex() != mc.Empty {
			co_child := model.CommentChild{}
			mc.Table(co_child.Table()).Where(bson.M{"_id": v.PID}).FindOne(&co_child)
			if co_child.ID.Hex() != mc.Empty {
				r_member := find_user_chlient.GetUserInfo(co_child.MID, false)
				if r_member.MID != 0 {
					child_one["u2"] = gin.H{
						"name":   r_member.UserName,
						"avatar": common.Avatar(r_member.Avatar),
						"mid":    r_member.MID,
					}
				}
			}
		}

		comment_child = append(comment_child, child_one)
	}

	result_json := c_code.V1GinSuccess(comment_child)
	c.JSON(200, result_json)
}
