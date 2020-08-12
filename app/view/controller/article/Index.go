package article

import (
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/app/view"
	"v2ex/model"
)

func Index(c *gin.Context) {
	_ht := defaultData(c)

	//获取最新的文章
	index_t := model.DataIndex{}

	_list := []model.DataIndex{}
	list := []gin.H{}
	mc.Table(index_t.Table()).Order(bson.M{"_id": -1}).Limit(10).Find(&_list)
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
	view.Render(c, "data/article_index", _ht)
}
