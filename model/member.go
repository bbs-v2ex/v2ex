package model

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type Member struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	MID        MIDTYPE            `json:"mid" bson:"mid"`
	MemberType int                `json:"member_type" bson:"member_type"`
	UserName   string             `json:"user_name" bson:"user_name"`
	Avatar     string             `json:"avatar" bson:"avatar"`
	More       MemberMore         `json:"-" bson:"-"`
}

func (t Member) Table() string {
	return "member"
}

func (t Member) IndexList() []mongo.IndexModel {

	return []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{"mid", bsonx.Int32(-1)},
			},
		},
		{
			Keys: bsonx.Doc{
				{"user_name", bsonx.Int32(-1)},
			},
		},
	}
}

func (t Member) EncryptionPassWord(password string) string {
	return c_code.Md516(c_code.Md532(password))
}

//获取用户的信息
func (t Member) GetUserInfo(mid MIDTYPE, find_more bool) (m Member) {
	mc.Table(t.Table()).Where(bson.M{"mid": mid}).FindOne(&m)
	if find_more {
		mc.Table(m.More.Table()).Where(bson.M{"_id": m.ID}).FindOne(&m.More)
	}
	return
}
