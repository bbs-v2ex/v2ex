package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CommentChild struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`

	MID int `json:"mid" bson:"mid"`
	//评论ID
	DID primitive.ObjectID `json:"cid" bson:"cid"`

	C2ID primitive.ObjectID `json:"c_2_id" bson:"c_2_id"`

	ZanLen int `json:"zan_len" bson:"zan_len"`
	//其他信息
	Info CommentText `json:"-" bson:"-"`
}

func (t CommentChild) Table() string {
	return "comment_child"
}
