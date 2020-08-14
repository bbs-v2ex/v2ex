package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type CommentQuestionRoot struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	//发布人ID
	MID MIDTYPE `json:"mid" bson:"mid"`
	//下级评论多少页
	RC int `json:"rc" bson:"rc"`
	//问题ID
	DID DIDTYPE `json:"did" bson:"did"`

	//多少人点赞
	ZanLen int `json:"zan_len" bson:"zan_len"`
	//其他信息
	Text CommentQuestionText `json:"-" bson:"-"`
}

func (t CommentQuestionRoot) Table() string {
	return "comment_question_root"
}
func (t CommentQuestionRoot) IndexList() []mongo.IndexModel {

	return []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{"mid", bsonx.Int32(-1)},
				{"did", bsonx.Int32(1)},
				{"zan_len", bsonx.Int32(-1)},
			},
			Options: &options.IndexOptions{},
		},
	}
}
