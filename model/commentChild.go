package model

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	return "comment_child"
}
