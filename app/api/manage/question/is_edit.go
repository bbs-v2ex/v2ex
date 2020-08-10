package question

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/app/api"
	"v2ex/model"
)

type _is_root_edit struct {
	DID model.DIDTYPE `json:"did" validate:"gt=0" comment:"文章ID"`
}

func is_root_edit(c *gin.Context) {
	_f := _is_root_edit{}
	c.BindJSON(&_f)
	//验证是否回答过此问题
	is_reply := model.CommentQuestionRoot{}
	err := mc.Table(is_reply.Table()).Where(bson.M{"mid": api.GetMID(c), "did": _f.DID}).FindOne(&is_reply)
	if err != nil {
		result_json := c_code.V1GinError(103, "查询失败")
		c.JSON(200, result_json)
		return
	}
	if is_reply.ID.Hex() != mc.Empty {
		result_json := c_code.V1GinSuccess(false)
		c.JSON(200, result_json)
		return
	} else {
		result_json := c_code.V1GinSuccess(true)
		c.JSON(200, result_json)
		return
	}
}
