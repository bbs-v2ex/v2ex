package model

//文章自增ID
const AutoIDArticle = 1

//问答自增ID 需要登录才可以回答
const AutoIDQA = 2

//讨论 需要登录才可以回答
const AutoIDDiscuss = 3

type AutoID struct {
	TableID int   `bson:"t_id"`
	ValID   int64 `bson:"val"`
}

func (t AutoID) Table() string {
	return "autoid"
}
