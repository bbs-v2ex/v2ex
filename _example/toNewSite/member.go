package main

import (
	"fmt"
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	bson2 "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"log"
	"strings"
	"v2ex"
	"v2ex/model"
)

func main() {
	v2ex.ConnectMongodb()
	file, err := ioutil.ReadFile("F:/studyseo/lao_data/member.json")
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

			"_id":         d2["_id"],
			"mid":         d2["mid"],
			"member_type": 0,
			"user_name":   d2["user_name"],
			"avatar":      "",
		}
		mc.Table(model.Member{}.Table()).Where(bson.M{"_id": d2["_id"]}).UpdateOneIsEmptyNewInsert(member)
		mc.Table(model.MemberMore{}.Table()).Where(bson.M{"_id": d2["_id"]}).UpdateOneIsEmptyNewInsert(model.MemberMore{
			ID: d2["_id"].(primitive.ObjectID),
		})
		fmt.Println(d)
	}
}
