package article

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/app/api"
	"v2ex/model"
)

type _comment_child_list struct {
	RID primitive.ObjectID `json:"rid" validate:"len" comment:"评论ID"`
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
	comment_child := []model.CommentChild{}
	mc.Table(model.CommentChild{}.Table()).Where(bson.M{"_id": _f.RID}).Find(&comment_child)
	c.JSON(200, comment_child)
}
