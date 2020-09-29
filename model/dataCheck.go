package model

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
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

func AddDataCheck(_check DataCheck) gin.H {
	_check.ID = primitive.NewObjectID()
	_check.Itime = time.Now()
	err := mc.Table(_check.Table()).Insert(_check)
	if err != nil {
		result_json := c_code.V1GinError(400, "添加审核表失败")

		return result_json
	}
	result_json := c_code.V1GinSuccess(200, "已进入后台审核,通过后会展示", UrlViewMemberConfig+"/data_check_view?id="+_check.ID.Hex())
	return result_json
}
