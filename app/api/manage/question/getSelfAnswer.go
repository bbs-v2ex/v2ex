package question

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/app/api"
	"v2ex/model"
)

type _get_self_answer struct {
	DID model.DIDTYPE `json:"did"`
}

func get_self_answer(c *gin.Context) {
	_f := _get_self_answer{}
	c.BindJSON(&_f)
	mid := api.GetMID(c)
	answer := model.CommentQuestionRoot{}
	err := mc.Table(answer.Table()).Where(bson.M{"mid": mid, "did": _f.DID}).FindOne(&answer)
	if err != nil || answer.ID.Hex() == mc.Empty {
		result_json := c_code.V1GinError(101, "查询失败")
		c.JSON(200, result_json)
		return
	}

	mc.Table(answer.Text.Table()).Where(bson.M{"_id": answer.ID}).FindOne(&answer.Text)
	text := api.RestorePicture(answer.Text.Text, "", answer.Text.Img)
	result_json := c_code.V1GinSuccess(gin.H{"text": text, "_id": answer.Text.ID})
	c.JSON(200, result_json)
	return
}
