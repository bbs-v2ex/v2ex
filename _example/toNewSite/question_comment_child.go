package main

import (
	"fmt"
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	bson2 "go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"strings"
	"v2ex"
	"v2ex/model"
)

func main() {
	v2ex.ConnectMongodb()
	file, err := ioutil.ReadFile("F:/studyseo/lao_data/comment_child.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, d := range strings.Split(strings.ReplaceAll(string(file), "\r", ""), "\n") {
		d2 := bson.M{}
		err := bson2.UnmarshalExtJSON([]byte(d), false, &d2)
		if err != nil {
			continue
		}
		member := bson.M{
			"mid":     d2["mid"],
			"rid":     d2["cid"],
			"pid":     d2["c_2_id"],
			"zan_len": d2["zan_len"],
		}
		mc.Table(model.CommentQuestionChild{}.Table()).Where(bson.M{"_id": d2["_id"]}).UpdateOneIsEmptyNewInsert(member)
		fmt.Println(d)
	}
}
