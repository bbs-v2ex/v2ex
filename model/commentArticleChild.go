package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type CommentChild struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`

	MID MIDTYPE `json:"mid" bson:"mid"`
	//评论 Root 主ID
	RID primitive.ObjectID `json:"rid" bson:"rid"`

	//父ID
	PID primitive.ObjectID `json:"pid" bson:"pid"`

	ZanLen int `json:"zan_len" bson:"zan_len"`
	//其他信息
	Text CommentText `json:"-" bson:"-"`
}

func (t CommentChild) Table() string {
	return "comment_article_child"
}
func (t CommentChild) IndexList() []mongo.IndexModel {

	return []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{"mid", bsonx.Int32(1)},
				{"rid", bsonx.Int32(1)},
				{"pid", bsonx.Int32(1)},
			},
			Options: &options.IndexOptions{},
		},
	}
}
