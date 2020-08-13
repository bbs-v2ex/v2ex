package controller

import (
	"fmt"
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"v2ex/app/api/member"
	"v2ex/app/nc"
	"v2ex/app/view"
	"v2ex/model"
)

func LastActivity(c *gin.Context) {
	seoconfig := nc.GetSeoConfig()
	xx := c.Query("xx")
	_ht := defaultData(c)
	rid, _ := primitive.ObjectIDFromHex(c.Param("rid"))

	list := member.ListDynamic(0, rid)
	_ht["dt"] = list
	_ht["next_link"] = ""

	if len(list) >= 10 {
		_ht["next_link"] = "/last_activity/" + list[len(list)-1].ID
	}
	if xx == "nohead" {
		html_content := view.RenderGetContent("_list/dongtai.html", _ht)
		_ht["content"] = c_code.CompressHtml(html_content)
		result_json := c_code.V1GinSuccess(html_content)
		result_json["next"] = _ht["next_link"]
		c.JSON(200, result_json)
		return
	}
	if len(list) <= 1 {
		c.Redirect(301, "/last_activity")
	}
	//加载热门问题
	t_30 := time.Now().AddDate(0, 0, -30)
	fmt.Println(t_30.String())
	index_t := model.DataIndex{}
	_article := []model.DataIndex{}
	article := []gin.H{}
	mc.Table(index_t.Table()).Where(bson.M{"ct": bson.M{"$gt": t_30.Unix()}, "d_type": model.DTYPEArticle}).Order(bson.M{"rc": -1, "_id": -1}).Limit(10).Find(&_article)

	for _, v := range _article {
		article = append(article, gin.H{
			"t": v.T,
			"u": model.UrlArticle(v),
		})
	}
	_ht["article"] = article
	//得到热门问题
	_question := []model.DataIndex{}
	question := []gin.H{}
	mc.Table(index_t.Table()).Where(bson.M{"ct": bson.M{"$gt": t_30.Unix()}, "d_type": model.DTYPEQuestion}).Order(bson.M{"rc": -1, "_id": -1}).Limit(10).Find(&_question)

	for _, v := range _question {
		question = append(question, gin.H{
			"t":  v.T,
			"u":  model.UrlQuestion(v),
			"rc": v.RC,
		})
	}
	_ht["question"] = question
	_ht["t"] = TitleJoin([]string{seoconfig.Activity.T})
	_ht["k"] = seoconfig.Activity.K
	_ht["d"] = DesJoin(seoconfig.Activity.D)
	view.Render(c, "_last_activity", _ht)
}
