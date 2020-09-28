package member

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/app/api"
	"v2ex/model"
)

func isCollect(c *gin.Context) {
	_f := _collect_add{}
	c.BindJSON(&_f)
	index := model.DataIndex{}
	mc.Table(index.Table()).Where(bson.M{"did": _f.DID}).FindOne(&index)
	if index.DID == 0 {
		result_json := c_code.V1GinError(101, "s数据丢失")
		c.JSON(200, result_json)
		return
	}

	mid := api.GetMID(c)
	collect := model.MemberCollect{}
	mc.Table(collect.Table()).Where(bson.M{"mid": mid, "did": index.DID}).FindOne(&collect)
	if collect.MID != 0 {
		result_json := c_code.V1GinSuccess(true)
		c.JSON(200, result_json)
		return
	}
	result_json := c_code.V1GinSuccess(false)
	c.JSON(200, result_json)
	return
}
