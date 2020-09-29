package nc

import (
	"errors"
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"v2ex/model"
)

func QuestionCommentChildAdd(rid primitive.ObjectID, pid primitive.ObjectID, txt string, mid model.MIDTYPE, l_time time.Time, child_id primitive.ObjectID) error {
	//检测评论数据是否存在
	comment_root := model.CommentQuestionRoot{}
	mc.Table(comment_root.Table()).Where(bson.M{"_id": rid}).FindOne(&comment_root)
	if comment_root.ID.Hex() == mc.Empty {
		return errors.New("请勿乱传参")
	}

	//通过 插入数据库
	comment_child := model.CommentQuestionChild{
		ID:     child_id,
		MID:    mid,
		RID:    rid,
		PID:    pid,
		ZanLen: 0,
	}
	err := mc.Table(comment_child.Table()).Insert(comment_child)
	if err != nil {
		return errors.New("写入child 表失败")
	}

	//写入数据存储表
	comment_text := model.CommentQuestionText{
		ID:          comment_child.ID,
		Text:        txt,
		Zan:         nil,
		Img:         nil,
		ReleaseTime: l_time,
	}
	//写进数据表中
	err = mc.Table(comment_text.Table()).Insert(comment_text)
	if err != nil {
		return errors.New("写入CommentText 表失败")
	}

	//主频率 RC 字段加一
	go func() {
		mc.Table(comment_root.Table()).Where(bson.M{"_id": rid}).FieldAddOrDel("rc", +1)
	}()
	return nil
}
