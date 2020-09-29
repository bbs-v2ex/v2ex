package nc

import (
	"errors"
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"v2ex/model"
)

func ArticleCommentRootAdd(did model.DIDTYPE, txt string, mid model.MIDTYPE, l_time time.Time, root_id primitive.ObjectID) error {

	//检测did 是否存在
	index := model.DataIndex{}
	mc.Table(index.Table()).Where(bson.M{"did": did, "d_type": model.DTYPEArticle}).FindOne(&index)
	if index.DID == 0 {
		return errors.New("数据丢失")
	}

	//写入评论索引表
	comment_root := model.CommentRoot{
		ID:     root_id,
		MID:    mid,
		RC:     0,
		DID:    index.DID,
		ZanLen: 0,
	}
	err := mc.Table(comment_root.Table()).Insert(comment_root)
	if err != nil {
		return errors.New("写入CommentRoot索引表失败")
	}

	//写入数据存储表
	comment_text := model.CommentText{
		ID:          comment_root.ID,
		Text:        txt,
		Zan:         nil,
		Img:         nil,
		ReleaseTime: l_time,
	}
	//写进数据表中
	err = mc.Table(comment_text.Table()).Insert(comment_text)

	if err != nil {
		return errors.New("写入CommentText表失败")
	}

	//评论字段加 1
	go func() {
		mc.Table(index.Table()).Where(bson.M{"_id": index.ID}).UpdateOne(bson.M{"rc": index.RC + 1, "ct": time.Now().Unix()})
		//添加进通知中心
		model.Movement(mid, index.MID).AddArticleCommentRoot(comment_root)
	}()
	return nil
}
