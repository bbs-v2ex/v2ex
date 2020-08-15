package edit

import (
	"fmt"
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/app/api"
	"v2ex/model"
)

func editpost(c *gin.Context) {
	_f := _edit_result{}
	err := c.Bind(&_f)
	if err != nil {
		result_json := c_code.V1GinError(100, err.Error())
		c.JSON(200, result_json)
		return
	}
	mid := api.GetMID(c)
	switch _f.Type {
	case "article":
		data_index := model.DataIndex{}
		mc.Table(data_index.Table()).Where(bson.M{"did": _f.DID, "d_type": model.DTYPEArticle, "mid": mid}).FindOne(&data_index)
		if data_index.MID != mid {
			result_json := c_code.V1GinError(101, "请勿乱传参")
			c.JSON(200, result_json)
			return
		}
		mc.Table(data_index.Table()).Where(bson.M{"_id": data_index.ID}).UpdateOne(bson.M{"t": _f.Title})
		content, imgs, err := api.SeparatePicture(_f.Content)
		if err != nil {
			result_json := c_code.V1GinError(102, "请勿乱传参")
			c.JSON(200, result_json)
			return
		}
		mc.Table(data_index.InfoArticle.Table()).Where(bson.M{"_id": data_index.ID}).UpdateOne(bson.M{"content": content, "imgs": imgs})

		result_json := c_code.V1GinSuccess("修改成功", "", fmt.Sprintf("/%s/%d", model.UrlTagArticle, _f.DID))
		c.JSON(200, result_json)
		return
	case "question":
		data_index := model.DataIndex{}
		mc.Table(data_index.Table()).Where(bson.M{"did": _f.DID, "d_type": model.DTYPEQuestion, "mid": mid}).FindOne(&data_index)
		if data_index.MID != mid {
			result_json := c_code.V1GinError(101, "请勿乱传参")
			c.JSON(200, result_json)
			return
		}
		mc.Table(data_index.Table()).Where(bson.M{"_id": data_index.ID}).UpdateOne(bson.M{"t": _f.Title})
		content, imgs, err := api.SeparatePicture(_f.Content)
		if err != nil {
			result_json := c_code.V1GinError(102, "请勿乱传参")
			c.JSON(200, result_json)
			return
		}
		mc.Table(data_index.InfoQuestion.Table()).Where(bson.M{"_id": data_index.ID}).UpdateOne(bson.M{"content": content, "imgs": imgs})

		result_json := c_code.V1GinSuccess("修改成功", "", fmt.Sprintf("/%s/%d", model.UrlTagQuestion, _f.DID))
		c.JSON(200, result_json)
		return

	case "question_answer":
		_html, _imgs, err2 := api.SeparatePicture(_f.Content)
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
		result_json := c_code.V1GinSuccess("修改成功", "", fmt.Sprintf("/%s/%d/%s/%s", model.UrlTagQuestion, _f.DID, model.UrlTagQuestionReply, answer.ID.Hex()))
		c.JSON(200, result_json)
		return
	}

}
