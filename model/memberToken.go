package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type MemberToken struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`

	//用户ID
	MID MIDTYPE `json:"mid" bson:"mid"`

	Token string `json:"token" bson:"token"`

	//最后活跃时间
	Expire string `json:"expire" bson:"expire"`
}

func (t MemberToken) Table() string {
	return "member_token"
}

func (t MemberToken) IndexList() []mongo.IndexModel {

	return []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{"token", bsonx.Int32(-1)},
			},
		},
	}
}
