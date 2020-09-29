package root_api

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/model"
)

type _data_del struct {
	ID primitive.ObjectID `json:"id"`
}

func dataDel(c *gin.Context) {
	_f := _data_del{}
	c.BindJSON(&_f)
	index := model.DataIndex{}
	mc.Table(index.Table()).Where(bson.M{"_id": _f.ID}).FindOne(&index)
	if index.ID.Hex() == mc.Empty {
		c.JSON(200, c_code.V1GinError(101, "未找到数据"))
		return
	}
	switch index.DTYPE {
	case model.DTYPEArticle:
		//查询所有的评论
		comment_root_list := []model.CommentRoot{}
		err := mc.Table(model.CommentRoot{}.Table()).Where(bson.M{"did": index.DID}).Find(&comment_root_list)
		if err != nil {
			c.JSON(200, c_code.V1GinError(101, "获取所有评论类别失败"))
			return
		}
		for _, v := range comment_root_list {
			child_list := []model.CommentChild{}
			mc.Table(model.CommentChild{}.Table()).Where(bson.M{"rid": v.ID}).Find(&child_list)
			for _, del := range child_list {
				mc.Table(model.CommentChild{}.Table()).Where(bson.M{"_id": del.ID}).DelOne()
				mc.Table(model.CommentText{}.Table()).Where(bson.M{"_id": del.ID}).DelOne()
			}

			count, err := mc.Table(model.CommentChild{}.Table()).Where(bson.M{"rid": v.ID}).Count()
			if err != nil {
				c.JSON(200, c_code.V1GinError(101, err.Error()))
				return
			}
			if count != 0 {
				c.JSON(200, c_code.V1GinError(101, "评论没清理完成,请重试"))
				return
			}
			mc.Table(model.CommentRoot{}.Table()).Where(bson.M{"_id": v.ID}).DelOne()

		}
		//检测是否还存在下级数据
		count, err := mc.Table(model.CommentRoot{}.Table()).Where(bson.M{"did": index.DID}).Count()
		if err != nil {
			c.JSON(200, c_code.V1GinError(101, "得到总数失败"))
			return
		}
		if count != 0 {
			c.JSON(200, c_code.V1GinError(101, "评论没清理完成,请重试"))
			return
		}
		//删除数据
		mc.Table(model.DataArticle{}.Table()).Where(bson.M{"_id": _f.ID}).DelOne()
		mc.Table(model.DataIndex{}.Table()).Where(bson.M{"_id": _f.ID}).DelOne()
		c.JSON(200, c_code.V1GinSuccess("删除成功"))
		return
	case model.DTYPEQuestion:
		//查询所有的评论
		comment_root_list := []model.CommentQuestionRoot{}
		err := mc.Table(model.CommentQuestionRoot{}.Table()).Where(bson.M{"did": index.DID}).Find(&comment_root_list)
		if err != nil {
			c.JSON(200, c_code.V1GinError(101, "获取所有评论类别失败"))
			return
		}
		for _, v := range comment_root_list {
			child_list := []model.CommentQuestionChild{}
			mc.Table(model.CommentQuestionChild{}.Table()).Where(bson.M{"rid": v.ID}).Find(&child_list)
			for _, del := range child_list {
				mc.Table(model.CommentQuestionChild{}.Table()).Where(bson.M{"_id": del.ID}).DelOne()
				mc.Table(model.CommentQuestionText{}.Table()).Where(bson.M{"_id": del.ID}).DelOne()
			}
			count, err := mc.Table(model.CommentQuestionChild{}.Table()).Where(bson.M{"rid": v.ID}).Count()
			if err != nil {
				c.JSON(200, c_code.V1GinError(101, err.Error()))
				return
			}
			if count != 0 {
				c.JSON(200, c_code.V1GinError(101, "评论没清理完成,请重试"))
				return
			}
			mc.Table(model.CommentQuestionRoot{}.Table()).Where(bson.M{"_id": v.ID}).DelOne()
		}
		//检测是否还存在下级数据
		count, err := mc.Table(model.CommentQuestionRoot{}.Table()).Where(bson.M{"did": index.DID}).Count()
		if err != nil {
			c.JSON(200, c_code.V1GinError(101, "得到总数失败"))
			return
		}
		if count != 0 {
			c.JSON(200, c_code.V1GinError(101, "评论没清理完成,请重试"))
			return
		}
		//删除数据
		mc.Table(model.DataQuestion{}.Table()).Where(bson.M{"_id": _f.ID}).DelOne()
		mc.Table(model.DataIndex{}.Table()).Where(bson.M{"_id": _f.ID}).DelOne()
		c.JSON(200, c_code.V1GinSuccess("删除成功"))
		return
	}
}
