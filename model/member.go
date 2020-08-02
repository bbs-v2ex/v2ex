package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Member struct {
	ID       primitive.ObjectID `bson:"id"`
	MID      MIDTYPE            `bson:"mid"`
	UserName string             `bson:"user_name"`
	Avatar   string             `bson:"avatar"`
	More     MemberMore         `bson:"-"`
}
