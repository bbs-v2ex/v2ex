package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type MemberCollect struct {
	ID  primitive.ObjectID `json:"_id" bson:"_id"`
	MID MIDTYPE            `json:"mid" bson:"mid"`
	DID DIDTYPE            `json:"did" bson:"did"`
}

func (t MemberCollect) Table() string {
	return "member_collect"
}
func (t MemberCollect) IndexList() []mongo.IndexModel {

	return []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{"mid", bsonx.Int32(-1)},
				{"did", bsonx.Int32(-1)},
			},
		},
	}
}
