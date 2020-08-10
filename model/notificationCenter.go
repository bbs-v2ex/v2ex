package model

import (
	"fmt"
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Notice
//Notification
type NoticeCenter struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	//发起人
	MID MIDTYPE `json:"mid" bson:"mid"`
	//接收人
	M2ID MIDTYPE     `json:"m2id" bson:"m2id"`
	Hash string      `json:"hash" bson:"hash"`
	Read bool        `json:"read" bson:"read"`
	Type int         `json:"type" bson:"type"`
	V    interface{} `json:"v" bson:"v"`
}

func (t NoticeCenter) Table() string {
	return "notice_center"
}

const (
	//发布文章
	NoticeArticleSend = 1

	//发布文章评论
	NoticeArticleCommentRoot = 2

	//文章点赞
	NoticeArticleCommentGood = 3
)

func Notice(mid, m2id MIDTYPE) NoticeCenter {
	return NoticeCenter{
		ID:   primitive.NewObjectID(),
		MID:  mid,
		M2ID: m2id,
	}
}

func (t NoticeCenter) verify() bool {
	if t.MID == t.M2ID {
		return false
	}
	return true
}

type articleSend struct {
	DID DIDTYPE `json:"did" bson:"did"`
}

//发布文章
func (t NoticeCenter) AddArticleSend(index DataIndex) {
	if !t.verify() {
		return
	}
	t.Type = NoticeArticleSend
	t.V = articleSend{
		DID: index.DID,
	}
	mc.Table(t.Table()).Insert(t)
}

type articleCommentRoot struct {
	RID primitive.ObjectID `json:"rid" bson:"rid"`
}

func (t NoticeCenter) AddArticleCommentRoot(index CommentRoot) {
	if !t.verify() {
		return
	}
	t.Type = NoticeArticleCommentRoot
	t.V = articleCommentRoot{
		RID: index.ID,
	}
	mc.Table(t.Table()).Insert(t)
}

type articleZan struct {
	RID primitive.ObjectID `json:"rid" bson:"rid"`
}

func (t NoticeCenter) AddArticleZan(index CommentRoot) {
	if !t.verify() {
		return
	}
	t.Type = NoticeArticleCommentGood
	t.V = articleZan{
		RID: index.ID,
	}
	where := bson.M{"mid": t.MID, "m2id": t.M2ID, "hash": index.ID.Hex()}
	//查询是否存在
	_t := NoticeCenter{}
	mc.Table(t.Table()).Where(where).FindOne(&_t)
	if _t.ID.Hex() != mc.Empty {
		return
	}
	t.Hash = index.ID.Hex()
	mc.Table(t.Table()).Insert(t)
}

func (t NoticeCenter) DelArticleZan(index CommentRoot) {
	if !t.verify() {
		return
	}
	where := bson.M{"mid": t.MID, "m2id": t.M2ID, "hash": index.ID.Hex()}
	fmt.Println(where)
	mc.Table(t.Table()).Where(where).DelOne()
}
