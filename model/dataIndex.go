package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

//数据索引
type DataIndex struct {
	ID  primitive.ObjectID `json:"_id" bson:"_id"`
	DID DIDTYPE            `json:"did" bson:"did"`
	//最后活跃时间 时间戳
	CT    int64   `json:"ct" bson:"ct"`
	DTYPE int     `json:"d_type" bson:"d_type"`
	MID   MIDTYPE `json:"mid" bson:"mid"`
	Show  int     `json:"show" bson:"show"`
	T     string  `json:"t" bson:"t"`
	//ReplyCount
	RC           int          `json:"rc" bson:"rc"`
	InfoQuestion DataQuestion `json:"-" bson:"-"`
	InfoArticle  DataArticle  `json:"-" bson:"-"`
}

func (t DataIndex) Table() string {
	return "data_index"
}

func (t DataIndex) IndexList() []mongo.IndexModel {

	return []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{"mid", bsonx.Int32(-1)},
				{"did", bsonx.Int32(1)},
				{"zan_len", bsonx.Int32(-1)},
			},
			Options: &options.IndexOptions{},
		},
		{
			Keys: bsonx.Doc{
				{"t", bsonx.Int32(-1)},
			},
		},
	}
}
