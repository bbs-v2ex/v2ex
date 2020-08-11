package member

import (
	"fmt"
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/app/api"
	"v2ex/model"
)

func collect_del(c *gin.Context) {
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
	where := bson.M{"mid": mid, "did": _f.DID}
	fmt.Println(where)
	mc.Table(model.MemberCollect{}.Table()).Where(where).DelOne()
	model.Movement(mid, index.MID).DelCollect(index)
	result_json := c_code.V1GinSuccess("删除成功")
	c.JSON(200, result_json)
	return
}
