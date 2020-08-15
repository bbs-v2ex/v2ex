package main

import (
	"fmt"
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex"
	"v2ex/model"
)

func main() {
	v2ex.ConnectMongodb()
	hex, _ := primitive.ObjectIDFromHex("5eb2ad8f808b9805d28c1989")
	a := model.DataQuestion{}
	err := mc.Table(a.Table()).Where(bson.M{"_id": hex}).FindOne(&a)
	fmt.Println(err)
	//5eb2ad8f808b9805d28c1989
}
