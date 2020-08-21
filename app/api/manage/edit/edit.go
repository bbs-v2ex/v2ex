package edit

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/app/api"
	"v2ex/model"
)

type _edit struct {
	DID  model.DIDTYPE `form:"did" json:"did"`
	Type string        `form:"type" json:"type"`
}

type _edit_result struct {
	_edit
	Title   string `json:"title"`
	Content string `json:"content"`
}

func edit(c *gin.Context) {
	_f := _edit{}
	err := c.Bind(&_f)
	if err != nil {
		result_json := c_code.V1GinError(100, err.Error())
		c.JSON(200, result_json)
		return
	}
	res := _edit_result{}
	res.DID = _f.DID
	res.Type = _f.Type

	mid := api.GetMID(c)
	switch _f.Type {
	case "article":
		data_index := model.DataIndex{}
		mc.Table(data_index.Table()).Where(bson.M{"did": _f.DID, "d_type": model.DTYPEArticle, "mid": mid}).FindOne(&data_index)
		if data_index.MID == 0 {
			result_json := c_code.V1GinError(101, "查询失败")
			c.JSON(200, result_json)
			return
		}
		res.Title = data_index.T
		mc.Table(data_index.InfoArticle.Table()).Where(bson.M{"_id": data_index.ID}).FindOne(&data_index.InfoArticle)
		if data_index.InfoArticle.Content == "" {
			result_json := c_code.V1GinError(102, "查询失败")
			c.JSON(200, result_json)
			return
		}
		res.Content = api.RestorePicture(data_index.InfoArticle.Content, "", data_index.InfoArticle.Imgs, "300")
		result_json := c_code.V1GinSuccess(res)
		c.JSON(200, result_json)
		return

	case "question":
		data_index := model.DataIndex{}
		mc.Table(data_index.Table()).Where(bson.M{"did": _f.DID, "d_type": model.DTYPEQuestion, "mid": mid}).FindOne(&data_index)
		if data_index.MID == 0 {
			result_json := c_code.V1GinError(101, "查询失败")
			c.JSON(200, result_json)
			return
		}
		res.Title = data_index.T
		mc.Table(data_index.InfoQuestion.Table()).Where(bson.M{"_id": data_index.ID}).FindOne(&data_index.InfoQuestion)
		if data_index.InfoQuestion.ID.Hex() == mc.Empty {
			result_json := c_code.V1GinError(102, "查询失败")
			c.JSON(200, result_json)
			return
		}
		res.Content = api.RestorePicture(data_index.InfoQuestion.Content, "", data_index.InfoQuestion.Imgs, "300")
		result_json := c_code.V1GinSuccess(res)
		c.JSON(200, result_json)
		return

	case "question_answer":
		answer := model.CommentQuestionRoot{}
		err := mc.Table(answer.Table()).Where(bson.M{"mid": mid, "did": _f.DID}).FindOne(&answer)
		if err != nil || answer.ID.Hex() == mc.Empty {
			result_json := c_code.V1GinError(101, "查询失败")
			c.JSON(200, result_json)
			return
		}

		mc.Table(answer.Text.Table()).Where(bson.M{"_id": answer.ID}).FindOne(&answer.Text)
		res.Content = api.RestorePicture(answer.Text.Text, "", answer.Text.Img, "300")

		result_json := c_code.V1GinSuccess(res)
		c.JSON(200, result_json)
		return
	}
}
