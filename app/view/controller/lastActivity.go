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
	"v2ex/app/view"
	"v2ex/model"
)

func LastActivity(c *gin.Context) {

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
		html_content := view.RenderGetContent("member/member_dz_activity.html", _ht)
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
	mc.Table(index_t.Table()).Where(bson.M{"ct": bson.M{"$gt": t_30.Unix()}, "d_type": model.DTYPEArticle}).Order(bson.M{"rc": -1}).Limit(10).Find(&_article)
	//如果不足10 条则以最新的数据填充
	if len(_article) != 10 {
		_a := []model.DataIndex{}
		mc.Table(index_t.Table()).Where(bson.M{"d_type": model.DTYPEArticle}).Order(bson.M{"_id": -1}).Limit(int64(10 - len(_article))).Find(&_a)
		_article = append(_article, _a...)
	}

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
	mc.Table(index_t.Table()).Where(bson.M{"ct": bson.M{"$gt": t_30.Unix()}, "d_type": model.DTYPEQuestion}).Order(bson.M{"rc": -1}).Limit(10).Find(&_question)

	if len(_question) != 10 {
		_a := []model.DataIndex{}
		mc.Table(index_t.Table()).Where(bson.M{"d_type": model.DTYPEQuestion}).Order(bson.M{"_id": -1}).Limit(int64(10 - len(_article))).Find(&_a)
		_question = append(_question, _a...)
	}

	for _, v := range _question {
		question = append(question, gin.H{
			"t":  v.T,
			"u":  model.UrlArticle(v),
			"rc": v.RC,
		})
	}
	_ht["question"] = question

	view.Render(c, "_last_activity", _ht)
}
