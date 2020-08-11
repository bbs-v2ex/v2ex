package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MemberCollect struct {
	ID  primitive.ObjectID `json:"_id" bson:"_id"`
	MID MIDTYPE            `json:"mid" bson:"mid"`
	DID DIDTYPE            `json:"did" bson:"did"`
}

func (t MemberCollect) Table() string {
	return "member_collect"
}
