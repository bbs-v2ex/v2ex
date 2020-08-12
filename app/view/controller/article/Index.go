package article

import (
	"fmt"
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"v2ex/app/nc"
	"v2ex/app/view"
	"v2ex/app/view/controller"
	"v2ex/model"
)

func Index(c *gin.Context) {
	seoconfig := nc.GetSeoConfig()
	_ht := defaultData(c)
	xx := c.Query("xx")
	//获取最新的文章
	index_t := model.DataIndex{}
	rid, _ := primitive.ObjectIDFromHex(c.Param("rid"))

	_list := []model.DataIndex{}
	list := []gin.H{}
	where := bson.M{"d_type": model.DTYPEArticle}
	if rid.Hex() != mc.Empty {
		where["_id"] = bson.M{"$lt": rid}
	}
	mc.Table(index_t.Table()).Where(where).Limit(10).Order(bson.M{"_id": -1}).Find(&_list)
	aids := []primitive.ObjectID{}
	for _, v := range _list {
		mc.Table(v.InfoArticle.Table()).Where(bson.M{"_id": v.ID}).FindOne(&v.InfoArticle)
		if v.InfoArticle.ID.Hex() == mc.Empty {
			continue
		}
		_one := gin.H{
			"t":   v.T,
			"u":   model.UrlArticle(v),
			"txt": model.DesSplit(v.InfoArticle.Content, 120),
			"img": "",
		}
		if len(v.InfoArticle.Imgs) >= 1 {
			_one["img"] = model.UrlImage(v.InfoArticle.Imgs[0])
		}
		list = append(list, _one)
		aids = append(aids, v.ID)
	}
	_ht["dt"] = list

	_ht["next_link"] = ""

	if len(_list) >= 10 {
		_ht["next_link"] = "/" + model.UrlTagArticle + "/_/l/" + _list[len(_list)-1].ID.Hex()
	}
	if xx == "nohead" {
		html_content := view.RenderGetContent("data/article_index_list.html", _ht)
		_ht["content"] = c_code.CompressHtml(html_content)
		result_json := c_code.V1GinSuccess(html_content)
		result_json["next"] = _ht["next_link"]
		c.JSON(200, result_json)
		return
	}
	if len(list) <= 1 {
		c.Redirect(301, "/last_activity")
	}
	//加载热门文章
	t_30 := time.Now().AddDate(0, 0, -30)
	_article := []model.DataIndex{}
	article := []gin.H{}
	mc.Table(index_t.Table()).Where(bson.M{"_id": bson.M{"$nin": aids}, "ct": bson.M{"$gt": t_30.Unix()}, "d_type": model.DTYPEArticle}).Order(bson.M{"rc": -1, "_id": -1}).Limit(10).Find(&_article)

	for _, v := range _article {
		article = append(article, gin.H{
			"t": v.T,
			"u": model.UrlArticle(v),
		})
		aids = append(aids, v.ID)
	}
	_ht["article"] = article

	//加载最新文章
	_article_new := []model.DataIndex{}
	article_new := []gin.H{}
	where = bson.M{"_id": bson.M{"$nin": aids}, "d_type": model.DTYPEArticle}
	mc.Table(index_t.Table()).Where(where).Order(bson.M{"_id": -1}).Limit(10).Find(&_article_new)

	ii := []string{}
	for _, v := range aids {
		ii = append(ii, v.Hex())
	}
	for _, v := range _article_new {
		if c_code.InArrayString(v.ID.Hex(), ii) {
			fmt.Println(v.ID.Hex())
		}
		article_new = append(article_new, gin.H{
			"t": v.T,
			"u": model.UrlArticle(v),
		})
	}
	_ht["article_new"] = article_new

	_ht["t"] = controller.TitleJoin([]string{seoconfig.Article.T})
	_ht["k"] = seoconfig.Article.K
	_ht["d"] = controller.DesJoin(seoconfig.Article.D)

	view.Render(c, "data/article_index", _ht)
}
