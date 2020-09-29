package nc

import (
	"errors"
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"v2ex/app/api"
	"v2ex/model"
)

func QuestionCommentRootAdd(txt string, did model.DIDTYPE, mid model.MIDTYPE, comment_root_id primitive.ObjectID) error {

	//分离数据
	_html, _imgs, err2 := api.SeparatePicture(txt)
	if err2 != nil {
		return errors.New("html解析错误")
	}

	//检测did 是否存在
	index := model.DataIndex{}
	mc.Table(index.Table()).Where(bson.M{"did": did, "d_type": model.DTYPEQuestion}).FindOne(&index)
	if index.DID == 0 {
		return errors.New("未找到该问题")
	}
	//通过 插入数据库
	comment_root := model.CommentQuestionRoot{
		ID:     comment_root_id,
		MID:    mid,
		RC:     0,
		DID:    index.DID,
		ZanLen: 0,
	}
	err := mc.Table(comment_root.Table()).Insert(comment_root)
	if err != nil {
		return errors.New("写入失败")
	}

	//写入数据存储表
	comment_text := model.CommentQuestionText{
		ID:          comment_root.ID,
		Text:        _html,
		Zan:         nil,
		Img:         _imgs,
		ReleaseTime: time.Now(),
	}
	//写进数据表中
	err = mc.Table(comment_text.Table()).Insert(comment_text)

	if err != nil {
		mc.Table(comment_root.Table()).Where(bson.M{"_id": comment_root.ID}).DelOne()
		return errors.New("写入失败")
	}

	//评论字段加 1
	mc.Table(index.Table()).Where(bson.M{"_id": index.ID}).UpdateOne(bson.M{"rc": index.RC + 1, "ct": time.Now().Unix()})

	model.Movement(mid, index.MID).AddQuestionCommentRoot(comment_root)
	return nil

}
