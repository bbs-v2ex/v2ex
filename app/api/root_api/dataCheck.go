package root_api

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/model"
)

type _dataCheck struct {
	ID   primitive.ObjectID `json:"_id"`
	Type string             `json:"type"`
}

func dataCheck(c *gin.Context) {
	_f := _dataCheck{}
	c.BindJSON(&_f)
	_data := model.DataCheck{}
	mc.Table(_data.Table()).Where(bson.M{"_id": _f.ID}).FindOne(&_data)
	switch _f.Type {
	//修改状态为通过
	case "pass":
		//先判断是什么类型
		switch _data.Type {
		//添加文章
		case model.DataCheckTypeArticleAdd:
			_article_add(c, &_f, &_data)
			return
		//修改文章
		case model.DataCheckTypeArticleEdit:
			_article_edit(c, &_f, &_data)
			return
		case model.DataCheckTypeArticleCommentRootAdd:
			_article_comment_root_add(c, &_f, &_data)
			return
		case model.DataCheckTypeArticleCommentChildAdd:
			_article_comment_child_add(c, &_f, &_data)
			return
		}
		break
	case "deny":
		//直接删除
		mc.Table(model.DataCheck{}.Table()).Where(bson.M{"_id": _f.ID}).DelOne()
		c.JSON(200, c_code.V1GinSuccess(200, "删除成功"))
	}
}
