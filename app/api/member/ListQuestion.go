package member

import (
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/model"
)

func ListQuestion(mid model.MIDTYPE, PID primitive.ObjectID) []gin.H {
	list := []model.DataIndex{}
	list_view := []gin.H{}
	where := bson.M{"mid": mid}
	mc.Table(model.DataIndex{}.Table()).Where(where).Order(bson.M{"_id": -1}).Limit(10).Find(&list)
	for _, v := range list {
		mc.Table(v.InfoQuestion.Table()).Where(bson.M{"_id": v.ID}).FindOne(&v.InfoQuestion)
		if v.InfoQuestion.ID.Hex() == mc.Empty {
			continue
		}
		_one := gin.H{
			"t":   v.T,
			"u":   model.UrlQuestion(v),
			"img": "",
			"txt": model.DesSplit(v.InfoQuestion.Content, 120),
		}

		if len(v.InfoQuestion.Imgs) >= 1 {
			_one["img"] = model.UrlImage(v.InfoQuestion.Imgs[0])
		}
		list_view = append(list_view, _one)
	}
	return list_view
}
