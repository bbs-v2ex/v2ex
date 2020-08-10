package main

import (
	"v2ex"
	"v2ex/model"
)

func main() {
	//链接数据库
	v2ex.ConnectMongodb()
	index := model.DataIndex{DID: 1110}
	model.Notice(789).SendArticle(index)

	//id := primitive.NewObjectID()
	//model.NotificationCenter{ID:  id, MID:  7, Read: false,}.SendArticle()
	//
	//mc.Table(model.NotificationCenter{}.Table()).Insert(inset)
	////
	//data := model.NotificationCenter{}
	//err := mc.Table(data.Table()).Where(bson.M{"_id": id}).FindOne(&data)
	//data.SendArticleGet()
	//fmt.Println(reflect.TypeOf(data.V).String())
	//fmt.Println(err)
	//fmt.Println(data)
}
