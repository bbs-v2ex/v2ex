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

func _article_add(c *gin.Context, _f *_dataCheck, _data *model.DataCheck) {

	article := nc.ArticleAdd(_data.MID, _data.D["title"].(string), _data.D["content"].(string), _data.Itime, true)
	article["u"] = ""
	if article["code"].(int) == 1 {
		mc.Table(model.DataCheck{}.Table()).Where(bson.M{"_id": _f.ID}).DelOne()
	}
	c.JSON(200, article)

}

func _article_edit(c *gin.Context, _f *_dataCheck, _data *model.DataCheck) {
	err := nc.ArticleEdit(_data.D["title"].(string), _data.D["content"].(string), _data.DID, _data.MID, _data.Itime, true)
	if err != nil {
		c.JSON(200, c_code.V1GinError(10000, err.Error()))
		return
	}
	mc.Table(model.DataCheck{}.Table()).Where(bson.M{"_id": _f.ID}).DelOne()
	c.JSON(200, c_code.V1GinSuccess("ok"))
	return
}

func _article_comment_root_add(c *gin.Context, _f *_dataCheck, _data *model.DataCheck) {
	err := nc.ArticleCommentRootAdd(_data.DID, _data.D["txt"].(string), _data.MID, _data.Itime, primitive.NewObjectID())
	if err != nil {
		c.JSON(200, c_code.V1GinError(10000, err.Error()))
		return
	}
	mc.Table(model.DataCheck{}.Table()).Where(bson.M{"_id": _f.ID}).DelOne()
	c.JSON(200, c_code.V1GinSuccess("ok"))
	return
}

func _article_comment_child_add(c *gin.Context, _f *_dataCheck, _data *model.DataCheck) {

	rid, _ := primitive.ObjectIDFromHex(_data.D["rid"].(string))
	pid, _ := primitive.ObjectIDFromHex(_data.D["pid"].(string))

	err := nc.ArticleCommentChildAdd(rid, pid, _data.D["txt"].(string), _data.MID, _data.Itime, primitive.NewObjectID())
	if err != nil {
		c.JSON(200, c_code.V1GinError(10000, err.Error()))
		return
	}
	mc.Table(model.DataCheck{}.Table()).Where(bson.M{"_id": _f.ID}).DelOne()
	c.JSON(200, c_code.V1GinSuccess("ok"))
	return
}
