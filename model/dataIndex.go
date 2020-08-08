package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//数据索引
type DataIndex struct {
	ID  primitive.ObjectID `json:"_id" bson:"_id"`
	DID DIDTYPE            `json:"did" bson:"did"`
	MID MIDTYPE            `json:"mid" bson:"mid"`
	T   string             `json:"t" bson:"t"`
	RC  int                `json:"rc" bson:"rc"`
}

func (t DataIndex) Table() string {
	return "data_index"
}
