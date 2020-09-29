package model

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DataCheck struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	//审核通过之后的ID
	Type int `json:"type" bson:"type"`
	//MID 操作人
	MID MIDTYPE `json:"mid" bson:"mid"`
	//DID 数据ID
	DID DIDTYPE `json:"did" bson:"did"`
	//数据
	D     gin.H     `json:"d" bson:"d"`
	Itime time.Time `json:"itime" bson:"itime"`
}

func (t DataCheck) Table() string {
	return "data_check"
}
