package article

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/app/api"
	"v2ex/model"
)

func list(c *gin.Context) {
	mid := api.GetMID(c)
	list := []model.DataIndex{}
	mc.Table(model.DataIndex{}.Table()).Where(bson.M{"d_type": model.DTYPEArticle, "mid": mid}).Limit(10).Find(&list)
	result_json := c_code.V1GinSuccess(list)
	c.JSON(200, result_json)
}
