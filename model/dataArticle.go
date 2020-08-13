package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//数据索引
type DataArticle struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Content string             `json:"content" bson:"content"`
	Imgs    []string           `json:"imgs" bson:"imgs"`
	//发布时间
	ReleaseTime time.Time `json:"release_time" bson:"release_time"`
	//调整时间
	ModifyTime    time.Time `json:"modify_time" bson:"modify_time"`
	LastReplyTime time.Time `json:"last_reply_time" bson:"last_reply_time"`
	LastReplyMID  int       `json:"last_reply_mid" bson:"last_reply_mid"`
	CommentSum    int       `json:"comment_sum" bson:"comment_sum"`
	CommentRoot   int       `json:"comment_root" bson:"comment_root"`
	//相关搜索
	RelatedTime time.Time `json:"related_time" bson:"related_time"`
	RelatedList []DIDTYPE `json:"related_list" bson:"related_list"`
}

func (t DataArticle) Table() string {
	return "data_article"
}
