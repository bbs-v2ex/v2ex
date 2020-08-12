package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//数据索引
type DataIndex struct {
	ID  primitive.ObjectID `json:"_id" bson:"_id"`
	DID DIDTYPE            `json:"did" bson:"did"`
	//最后活跃时间 时间戳
	CT           int64        `json:"ct" bson:"ct"`
	DTYPE        int          `json:"d_type" bson:"d_type"`
	MID          MIDTYPE      `json:"mid" bson:"mid"`
	Show         int          `json:"show" bson:"show"`
	T            string       `json:"t" bson:"t"`
	RC           int          `json:"rc" bson:"rc"`
	InfoQuestion DataQuestion `json:"-" bson:"-"`
	InfoArticle  DataArticle  `json:"-" bson:"-"`
}

func (t DataIndex) Table() string {
	return "data_index"
}
