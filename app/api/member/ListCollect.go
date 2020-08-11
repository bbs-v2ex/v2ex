package member

import (
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/model"
)

func ListCollect(mid model.MIDTYPE, PID primitive.ObjectID) []gin.H {
	list := []model.MemberCollect{}
	list_view := []gin.H{}
	where := bson.M{"mid": mid}
	mc.Table(model.MemberCollect{}.Table()).Where(where).Order(bson.M{"_id": -1}).Limit(10).Find(&list)
	for _, v := range list {
		index := model.DataIndex{}
		mc.Table(index.Table()).Where(bson.M{"did": v.DID}).FindOne(&index)
		if index.ID.Hex() == mc.Empty {
			continue
		}
		u := ""
		if index.DTYPE == model.DTYPEArticle {
			u = model.UrlArticle(index)
		}

		if index.DTYPE == model.DTYPEQuestion {
			u = model.UrlQuestion(index)
		}

		_one := gin.H{
			"t": index.T,
			"u": u,
		}

		list_view = append(list_view, _one)
	}
	return list_view
}
