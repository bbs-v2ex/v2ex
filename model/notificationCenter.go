package model

import (
	"github.com/123456/c_code/mc"
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
	Read bool        `json:"read" bson:"read"`
	Type int         `json:"type" bson:"type"`
	V    interface{} `json:"v" bson:"v"`
}

type sendArticle struct {
	DID DIDTYPE `json:"did" bson:"did"`
}

const (
	NoticeSendArticle = 1
)

func Notice(mid MIDTYPE) NoticeCenter {
	return NoticeCenter{
		ID:  primitive.NewObjectID(),
		MID: mid,
	}
}
func (t NoticeCenter) SendArticle(index DataIndex) {
	t.Type = NoticeSendArticle
	t.V = sendArticle{
		DID: index.DID,
	}
	mc.Table(t.Table()).Insert(t)
}

func (t NoticeCenter) Table() string {
	return "notice_center"
}
