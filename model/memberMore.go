package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MemberMore struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	MID      MIDTYPE            `json:"mid" bson:"mid"`
	PassWord string             `json:"pass_word" bson:"pass_word"`
	Des      string             `json:"des" bson:"des"`
}

func (t MemberMore) Table() string {
	return "member_more"
}
