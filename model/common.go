package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CommentRoot struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	//发布人ID
	MID int `json:"mid" bson:"mid"`
	//下级评论多少页
	ReplyCount int `json:"reply_count" bson:"reply_count"`
	//问题ID
	QID primitive.ObjectID `json:"qid" bson:"qid"`
	//楼层编号
	ZanLen int `json:"zan_len" bson:"zan_len"`
	//其他信息
}

func (t CommentRoot) Table() string {
	return "comment_root"
}
