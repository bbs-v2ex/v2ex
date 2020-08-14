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
	file, err := ioutil.ReadFile("F:/studyseo/lao_data/question_more.json")
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
			"_id":             d2["_id"],
			"content":         d2["content"],
			"imgs":            d2["img"],
			"release_time":    d2["release_time"],
			"modify_time":     d2["modify_time"],
			"last_reply_time": d2["last_reply_time"],
			"last_reply_mid":  d2["last_reply_mid"],
			"comment_sum":     d2["comment_sum"],
			"comment_root":    d2["comment_sum"],
		}
		mc.Table(model.DataQuestion{}.Table()).Where(bson.M{"_id": d2["_id"]}).UpdateOneIsEmptyNewInsert(member)
		fmt.Println(d)
	}
}
