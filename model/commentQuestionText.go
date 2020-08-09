package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CommentQuestionText struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	//评论得文本
	Text string `json:"text" bson:"text"`
	//为此回答点赞得人
	Zan []MIDTYPE `json:"zan" bson:"zan"`
	Img []string  `json:"img" bson:"img"`
	//发布时间
	ReleaseTime time.Time `json:"release_time" bson:"release_time"`
}

func (t CommentQuestionText) Table() string {
	return "comment_question_text"
}
