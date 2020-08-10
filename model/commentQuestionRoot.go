package model

import "go.mongodb.org/mongo-driver/bson/primitive"

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