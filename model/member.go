package model

import (
	"github.com/123456/c_code"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Member struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	MID      MIDTYPE            `json:"mid" bson:"mid"`
	UserName string             `json:"user_name" bson:"user_name"`
	Avatar   string             `json:"avatar" bson:"avatar"`
	More     MemberMore         `json:"-" bson:"-"`
}

func (t Member) Table() string {
	return "member"
}

func (t Member) EncryptionPassWord(password string) string {
	return c_code.Md516(c_code.Md532(password))
}
