package question

import (
	"fmt"
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/app/api"
	"v2ex/model"
)

type _edit_self_answer struct {
	DID  model.DIDTYPE `json:"did" validate:"gt=0" comment:"文章ID"`
	Txt  string        `json:"txt" validate:"min=10,max=1000" comment:"数据"`
	Text string        `validate:"min=10,max=1000" comment:"数据"`
}

func edit_self_answer(c *gin.Context) {
	_f := _edit_self_answer{}
	c.BindJSON(&_f)
	_f.Text = c_code.RemoveHtmlTag(_f.Txt)
	validator := api.VerifyValidator(_f)

	if validator != "" {
		result_json := c_code.V1GinError(101, validator)
		c.JSON(200, result_json)
		return
	}

	//分离数据
	_html, _imgs, err2 := api.SeparatePicture(_f.Txt)
	if err2 != nil {
		result_json := c_code.V1GinError(102, "html解析错误")
		c.JSON(200, result_json)
		return
	}

	mid := api.GetMID(c)
	answer := model.CommentQuestionRoot{}
	err := mc.Table(answer.Table()).Where(bson.M{"mid": mid, "did": _f.DID}).FindOne(&answer)
	if err != nil || answer.ID.Hex() == mc.Empty {
		result_json := c_code.V1GinError(103, "查询失败")
		c.JSON(200, result_json)
		return
	}
	update := bson.M{
		"text": _html,
		"img":  _imgs,
	}
	err = mc.Table(answer.Text.Table()).Where(bson.M{"_id": answer.ID}).UpdateOne(update)
	if err != nil {
		result_json := c_code.V1GinError(104, "修改失败")
		c.JSON(200, result_json)
		return
	}
	result_json := c_code.V1GinSuccess("修改成功", "", fmt.Sprintf("/q/%d/answer/%s", _f.DID, answer.ID.Hex()))
	c.JSON(200, result_json)
	return
}
