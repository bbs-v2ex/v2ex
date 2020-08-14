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
	file, err := ioutil.ReadFile("F:/studyseo/lao_data/question.json")
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
			"_id":    d2["_id"],
			"did":    0,
			"ct":     d2["rt"],
			"d_type": model.DTYPEQuestion,
			"mid":    d2["mid"],
			"t":      d2["t"],
			"rc":     d2["reply_sum"],
		}
		mc.Table(model.DataIndex{}.Table()).Where(bson.M{"_id": d2["_id"]}).UpdateOneIsEmptyNewInsert(member)
		fmt.Println(d)
	}
}
