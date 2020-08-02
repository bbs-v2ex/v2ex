package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//文章自增ID
const AutoIDArticle = 1

//问答自增ID 需要登录才可以回答
const AutoIDQA = 2

//讨论 需要登录才可以回答
const AutoIDDiscuss = 3

const AutoIDMember = 4

type AutoID struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	TableID int                `json:"t_id" bson:"t_id"`
	ValID   int64              `json:"val" bson:"val"`
}

func (t AutoID) Table() string {
	return "autoid"
}
