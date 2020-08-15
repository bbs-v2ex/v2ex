package main

import (
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	"v2ex"
	"v2ex/model"
)

func main() {
	v2ex.ConnectMongodb()
	list := []model.Member{}
	mc.Table(model.Member{}.Table()).Where(bson.M{}).Find(&list)
	for _, v := range list {
		mc.Table(v.Table()).Where(bson.M{"_id": v.ID}).UpdateOne(bson.M{"is_user": false})
	}
}
