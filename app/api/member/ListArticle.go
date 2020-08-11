package member

import (
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/model"
)

func ListArticle(mid model.MIDTYPE, PID primitive.ObjectID) []gin.H {
	list := []model.DataIndex{}
	list_view := []gin.H{}
	where := bson.M{"mid": mid}
	mc.Table(model.DataIndex{}.Table()).Where(where).Order(bson.M{"_id": -1}).Limit(10).Find(&list)
	for _, v := range list {
		mc.Table(v.InfoArticle.Table()).Where(bson.M{"_id": v.ID}).FindOne(&v.InfoArticle)
		if v.InfoArticle.ID.Hex() == mc.Empty {
			continue
		}
		_one := gin.H{
			"t":   v.T,
			"u":   model.UrlQuestion(v),
			"img": "",
			"txt": model.DesSplit(v.InfoArticle.Content, 120),
		}

		if len(v.InfoArticle.Imgs) >= 1 {
			_one["img"] = model.UrlImage(v.InfoArticle.Imgs[0])
		}
		list_view = append(list_view, _one)
	}
	return list_view
}
