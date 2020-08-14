package main

import (
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	"v2ex"
	"v2ex/model"
	"v2ex/until"
)

func main() {
	v2ex.ConnectMongodb()
	list := []model.Member{}
	mc.Table(model.Member{}.Table()).Where(bson.M{"avatar": ""}).Find(&list)
	for _, v := range list {

		mc.Table(v.Table()).Where(bson.M{"_id": v.ID}).UpdateOne(bson.M{"avatar": until.RandomAvatar()})

	}
}
