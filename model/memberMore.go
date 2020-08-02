package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MemberMore struct {
	ID       primitive.ObjectID `bson:"id"`
	MID      MIDTYPE
	PassWord string `bson:"pass_word"`
}

func (t MemberMore) Table() string {
	return "member_more"
}
