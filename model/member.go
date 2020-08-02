package model

import (
	"github.com/123456/c_code"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Member struct {
	ID       primitive.ObjectID `bson:"id"`
	MID      MIDTYPE            `bson:"mid"`
	UserName string             `bson:"user_name"`
	Avatar   string             `bson:"avatar"`
	More     MemberMore         `bson:"-"`
}

func (t Member) Table() string {
	return "member"
}

func (t Member) EncryptionPassWord(password string) string {
	return c_code.Md516(c_code.Md532(password))
}
