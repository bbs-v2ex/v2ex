package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MemberMore struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	PassWord    string             `json:"pass_word" bson:"pass_word"`
	Des         string             `json:"des" bson:"des"`
	DesDetailed string             `json:"des_detailed" bson:"des_detailed"`
}

func (t MemberMore) Table() string {
	return "member_more"
}
