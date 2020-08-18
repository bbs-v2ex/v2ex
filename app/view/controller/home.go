package controller

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/app/view"
	"v2ex/model"
	"v2ex/until"
)

func Home(c *gin.Context) {
	_ht := defaultData(c)
	xx := c.Query("xx")
	//获取最新的文章
	index_t := model.DataIndex{}
	rid, _ := primitive.ObjectIDFromHex(c.Param("rid"))

	_list := []model.DataIndex{}
	list := []gin.H{}
	where := bson.M{}
	if rid.Hex() != mc.Empty {
		where["_id"] = bson.M{"$lt": rid}
	}
	mc.Table(index_t.Table()).Where(where).Limit(10).Order(bson.M{"_id": -1}).Find(&_list)
	aids := []primitive.ObjectID{}
	for _, v := range _list {
		_one := gin.H{}
		switch v.DTYPE {
		case model.DTYPEArticle:
			mc.Table(v.InfoArticle.Table()).Where(bson.M{"_id": v.ID}).FindOne(&v.InfoArticle)
			if v.InfoArticle.ID.Hex() == mc.Empty {
				continue
			}
			_one = gin.H{
				"t":   v.T,
				"u":   model.UrlArticle(v),
				"txt": model.DesSplit(v.InfoArticle.Content, 120),
				"img": "",
			}
			if len(v.InfoArticle.Imgs) >= 1 {
				_one["img"] = model.UrlImage(v.InfoArticle.Imgs[0])
			}
			_one["tag"] = []gin.H{
				{
					"class": "i-tag-blue",
					"t":     "文章",
				},
			}
			break
		case model.DTYPEQuestion:
			mc.Table(v.InfoQuestion.Table()).Where(bson.M{"_id": v.ID}).FindOne(&v.InfoQuestion)
			if v.InfoQuestion.ID.Hex() == mc.Empty {
				continue
			}
			_one = gin.H{
				"t":   v.T,
				"u":   model.UrlQuestion(v),
				"txt": model.DesSplit(v.InfoQuestion.Content, 120),
				"img": "",
			}
			_one["tag"] = []gin.H{
				{
					"class": "i-tag-red",
					"t":     "问题",
				},
			}
			if len(v.InfoQuestion.Imgs) >= 1 {
				_one["img"] = model.UrlImage(v.InfoQuestion.Imgs[0])
			}
			break
		}

		list = append(list, _one)
		aids = append(aids, v.ID)
	}
	_ht["dt"] = list

	_ht["next_link"] = ""

	if len(_list) >= 10 {
		_ht["next_link"] = "/l/" + _list[len(_list)-1].ID.Hex()
	}
	if xx == "nohead" {
		html_content := view.RenderGetContent("_list/home.html", _ht)
		_ht["content"] = c_code.CompressHtml(html_content)
		result_json := c_code.V1GinSuccess(html_content)
		result_json["next"] = _ht["next_link"]
		c.JSON(200, result_json)
		return
	}
	if len(list) <= 1 {
		c.Redirect(301, "/")
	}
	_ht["dt"] = list

	//渲染最热数据 获取数据设置的时差
	t := until.DataTimeDifferenceIndexHome()
	//得到最热的文章
	_article := []model.DataIndex{}
	article := []gin.H{}

	hot_where := bson.M{"_id": bson.M{"$nin": aids}, "ct": bson.M{"$gt": t.Unix()}}
	hot_where["d_type"] = model.DTYPEArticle
	mc.Table(index_t.Table()).Where(hot_where).Order(bson.M{"rc": -1, "_id": -1}).Limit(10).Find(&_article)
	for _, v := range _article {
		article = append(article, gin.H{
			"t": v.T,
			"u": model.UrlArticle(v),
		})
		aids = append(aids, v.ID)
	}
	_ht["article"] = article
	//得到最热的提问
	_q := []model.DataIndex{}
	q := []gin.H{}
	hot_where["d_type"] = model.DTYPEQuestion
	mc.Table(index_t.Table()).Where(hot_where).Order(bson.M{"rc": -1, "_id": -1}).Limit(10).Find(&_q)
	for _, v := range _q {
		q = append(q, gin.H{
			"t": v.T,
			"u": model.UrlQuestion(v),
		})
		aids = append(aids, v.ID)
	}
	_ht["question"] = q

	//得到活跃会员
	members := model.MovementCenter{}.GetHuoYueMID(10)
	_ht["members"] = members

	view.Render(c, "index", _ht)
}
