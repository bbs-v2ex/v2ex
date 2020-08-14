package main

import (
	"fmt"
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	"strings"
	"v2ex"
	"v2ex/model"
)

func main() {
	v2ex.ConnectMongodb()
	list := []model.DataQuestion{}
	mc.Table(model.DataQuestion{}.Table()).Where(bson.M{}).Find(&list)
	for index, v := range list {
		fmt.Println(index, v.ID.Hex())
		if len(v.Imgs) == 0 {
			continue
		}
		img_list := []string{}
		for _, im := range v.Imgs {
			img_list = append(img_list, strings.ReplaceAll(im, "/static/data_img/", "/_old_site/"))
		}
		fmt.Println(v.ID.Hex())
		mc.Table(model.DataQuestion{}.Table()).Where(bson.M{"_id": v.ID}).UpdateOne(bson.M{"imgs": img_list})
	}
}
