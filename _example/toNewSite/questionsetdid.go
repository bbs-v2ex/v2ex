package main

import (
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	"log"
	"v2ex"
	"v2ex/model"
)

func main() {
	v2ex.ConnectMongodb()
	list := []model.DataIndex{}
	mc.Table(model.DataIndex{}.Table()).Where(bson.M{"did": 0}).Find(&list)
	for _, v := range list {
		qaid, err := model.AutoID{}.QAID()
		if err != nil {
			log.Fatal("错误")
		}
		err = mc.Table(v.Table()).Where(bson.M{"_id": v.ID}).UpdateOne(bson.M{"did": qaid})
		if err != nil {
			log.Fatal("错误")
		}
		model.AutoID{}.QAAdd()
	}
}
