package main

import (
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	bson2 "go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
	"v2ex"
	"v2ex/model"
)

func main() {
	v2ex.ConnectMongodb()
	file, err := ioutil.ReadFile("F:/studyseo/lao_data/member_more.json")
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
			"_id":           d2["_id"],
			"register_time": d2["register"],
			"des":           d2["des"],
		}
		mc.Table(model.MemberMore{}.Table()).Where(bson.M{"_id": d2["_id"]}).UpdateOneIsEmptyNewInsert(member)
		if _, ok := d2["avatar"]; ok {

			if reflect.TypeOf(d2["avatar"]).String() == "string" {
				if strings.HasPrefix(d2["avatar"].(string), "head-") {
					mc.Table(model.Member{}.Table()).Where(bson.M{"_id": d2["_id"]}).UpdateOneIsEmptyNewInsert(bson.M{"avatar": "/_avatar/" + d2["avatar"].(string)})
				}
			}
		}
	}
}
