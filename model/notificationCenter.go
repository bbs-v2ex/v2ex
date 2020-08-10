package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//Notice
//Notification
type NoticeCenter struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	MID  MIDTYPE            `json:"mid" bson:"mid"`
	Read bool               `json:"read" bson:"read"`
	Type int                `json:"type" bson:"type"`
	V    interface{}        `json:"v" bson:"v"`
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
func (t *NoticeCenter) SendArticle(index DataIndex) {
	t.Type = NoticeSendArticle
	t.V = sendArticle{
		DID: index.DID,
	}
}

func (t NoticeCenter) Table() string {
	return "notice_center"
}
