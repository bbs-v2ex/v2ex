package main

import (
	"fmt"
	"github.com/123456/c_code/mc"
	"github.com/PuerkitoBio/goquery"
	"github.com/globalsign/mgo/bson"
	"net/url"
	"strings"
	"v2ex"
	"v2ex/app/api"
	"v2ex/model"
)

func main() {
	v2ex.ConnectMongodb()
	list := []model.DataQuestion{}
	mc.Table(model.DataQuestion{}.Table()).Projection(bson.M{"imgs": 1, "content": 1}).Where(bson.M{}).Find(&list)
	for index, v := range list {
		fmt.Println(index, v.ID.Hex())
		//if len(v.Imgs) == 0 {
		//	continue
		//}

		//先还原图片
		content := api.RestorePicture(v.Content, "", v.Imgs)
		reader, err := goquery.NewDocumentFromReader(strings.NewReader(content))
		if err != nil {
			continue
		}
		if v.ID.Hex() == "5ed7b4e98831976c754135c2" {
			fmt.Println(123)
		}
		img_list := []string{}
		reader.Find("img").Each(func(i int, selection *goquery.Selection) {
			attr, _ := selection.Attr("src")
			if strings.HasPrefix(attr, "https://studyseo.net") {
				selection.SetAttr("src", "https://127.0.0.1/"+strings.ReplaceAll(attr, "https://studyseo.net", ""))
			}
		})
		reader.Find("img").Each(func(i int, selection *goquery.Selection) {
			attr, _ := selection.Attr("src")
			selection.ReplaceWithHtml("{{img}}")
			parse, err2 := url.Parse(attr)
			if err2 != nil {
				return
			}
			img_list = append(img_list, strings.ReplaceAll(parse.Path, "/static/data_img/", "/_old_site/"))
		})

		if len(img_list) > 0 {
			content, _ = reader.Find("body").Html()
			fmt.Println(content)
			mc.Table(model.DataQuestion{}.Table()).Where(bson.M{"_id": v.ID}).UpdateOne(bson.M{"imgs": img_list, "content": content})
		}

	}
}
