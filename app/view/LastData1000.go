package view

import (
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/config"
	"v2ex/model"
)

func LastData1000(c *gin.Context) {
	_con := config.GetConfig()
	site := _con.Run.SiteMapUrlPreFix
	list := []model.DataIndex{}
	mc.Table(model.DataIndex{}.Table()).Order(bson.M{"_id": -1}).Limit(1000).Find(&list)
	u := []string{}
	for _, v := range list {
		if v.DTYPE == model.DTYPEArticle {
			u = append(u, site+model.UrlArticle(v))
		} else {
			u = append(u, site+model.UrlQuestion(v))
		}
	}
	c.JSON(200, u)
}
