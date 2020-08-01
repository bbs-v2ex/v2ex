package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Member struct {
	ID       primitive.ObjectID
	MID      MIDTYPE
	UserName string
	Avatar   string
}
