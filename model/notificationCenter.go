package model

import (
	"fmt"
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//Notice
//Notification
type MovementCenter struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	//发起人
	MID MIDTYPE `json:"mid" bson:"mid"`
	//接收人
	M2ID        MIDTYPE     `json:"m2id" bson:"m2id"`
	Hash        string      `json:"hash" bson:"hash"`
	Read        bool        `json:"read" bson:"read"`
	Type        int         `json:"type" bson:"type"`
	V           interface{} `json:"v" bson:"v"`
	ReleaseTime time.Time   `json:"release_time" bson:"release_time"`
}

func (t MovementCenter) Table() string {
	return "movement_center"
}

const (
	//发布文章
	MovementArticleSend = 1

	//发布文章评论
	MovementArticleCommentRoot = 2

	//文章点赞
	MovementArticleCommentGood = 3
)

func Movement(mid, m2id MIDTYPE) MovementCenter {
	return MovementCenter{
		ID:          primitive.NewObjectID(),
		MID:         mid,
		M2ID:        m2id,
		ReleaseTime: time.Now(),
	}
}

func (t MovementCenter) verify() bool {
	if t.MID == t.M2ID {
		return false
	}
	return true
}

type articleSend struct {
	DID DIDTYPE `json:"did" bson:"did"`
}

//发布文章
func (t MovementCenter) AddArticleSend(index DataIndex) {
	if !t.verify() {
		return
	}
	t.Type = MovementArticleSend
	t.V = articleSend{
		DID: index.DID,
	}
	mc.Table(t.Table()).Insert(t)
}

type articleCommentRoot struct {
	RID primitive.ObjectID `json:"rid" bson:"rid"`
}

func (t MovementCenter) AddArticleCommentRoot(index CommentRoot) {
	if !t.verify() {
		return
	}
	t.Type = MovementArticleCommentRoot
	t.V = articleCommentRoot{
		RID: index.ID,
	}
	mc.Table(t.Table()).Insert(t)
}

type articleZan struct {
	RID primitive.ObjectID `json:"rid" bson:"rid"`
}

func (t MovementCenter) AddArticleZan(index CommentRoot) {
	if !t.verify() {
		return
	}
	t.Type = MovementArticleCommentGood
	t.V = articleZan{
		RID: index.ID,
	}
	where := bson.M{"mid": t.MID, "m2id": t.M2ID, "hash": index.ID.Hex()}
	//查询是否存在
	_t := MovementCenter{}
	mc.Table(t.Table()).Where(where).FindOne(&_t)
	if _t.ID.Hex() != mc.Empty {
		return
	}
	t.Hash = index.ID.Hex()
	mc.Table(t.Table()).Insert(t)
}

func (t MovementCenter) DelArticleZan(index CommentRoot) {
	if !t.verify() {
		return
	}
	where := bson.M{"mid": t.MID, "m2id": t.M2ID, "hash": index.ID.Hex()}
	fmt.Println(where)
	mc.Table(t.Table()).Where(where).DelOne()
}
