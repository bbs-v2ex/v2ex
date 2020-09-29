package root_api

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/app/nc"
	"v2ex/model"
)

func _question_add(c *gin.Context, _f *_dataCheck, _data *model.DataCheck) {

	err, index := nc.QuestionAdd(_data.D["title"].(string), _data.D["content"].(string), _data.MID, _data.Itime, true)
	if err != nil {
		c.JSON(200, c_code.V1GinError(1000, err.Error()))
		return
	}
	success := c_code.V1GinSuccess("", "", model.UrlQuestion(index))
	mc.Table(model.DataCheck{}.Table()).Where(bson.M{"_id": _f.ID}).DelOne()
	c.JSON(200, success)

}

func _question_edit(c *gin.Context, _f *_dataCheck, _data *model.DataCheck) {
	err := nc.QuestionEdit(_data.D["title"].(string), _data.D["content"].(string), _data.DID, _data.MID, _data.Itime, true)
	if err != nil {
		c.JSON(200, c_code.V1GinError(10000, err.Error()))
		return
	}
	mc.Table(model.DataCheck{}.Table()).Where(bson.M{"_id": _f.ID}).DelOne()
	c.JSON(200, c_code.V1GinSuccess("ok"))
	return
}

func _question_comment_root_add(c *gin.Context, _f *_dataCheck, _data *model.DataCheck) {

	err := nc.QuestionCommentRootAdd(_data.D["txt"].(string), _data.DID, _data.MID, primitive.NewObjectID())
	if err != nil {
		c.JSON(200, c_code.V1GinError(10000, err.Error()))
		return
	}
	mc.Table(model.DataCheck{}.Table()).Where(bson.M{"_id": _f.ID}).DelOne()
	c.JSON(200, c_code.V1GinSuccess("ok"))
	return
}
func _question_comment_root_edit(c *gin.Context, _f *_dataCheck, _data *model.DataCheck) {

	err := nc.QuestionCommentRootAdd(_data.D["txt"].(string), _data.DID, _data.MID, primitive.NewObjectID())
	if err != nil {
		c.JSON(200, c_code.V1GinError(10000, err.Error()))
		return
	}
	mc.Table(model.DataCheck{}.Table()).Where(bson.M{"_id": _f.ID}).DelOne()
	c.JSON(200, c_code.V1GinSuccess("ok"))
	return
}
