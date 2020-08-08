package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CommentRoot struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	//发布人ID
	MID MIDTYPE `json:"mid" bson:"mid"`
	//下级评论多少页
	ReplyCount int `json:"reply_count" bson:"reply_count"`
	//问题ID
	DID DIDTYPE `json:"did" bson:"did"`

	//多少人点赞
	ZanLen int `json:"zan_len" bson:"zan_len"`
	//其他信息
}

func (t CommentRoot) Table() string {
	return "comment_root"
}
