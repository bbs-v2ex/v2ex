package member

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/app/api"
	"v2ex/model"
)

type _collect_add struct {
	DID model.DIDTYPE `json:"did"`
}

func collectAdd(c *gin.Context) {
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
	if index.MID == mid {
		result_json := c_code.V1GinError(102, "自己的文章无法收藏")
		c.JSON(200, result_json)
		return
	}

	//查询是否已收藏
	collect := model.MemberCollect{}

	mc.Table(collect.Table()).Where(bson.M{"mid": mid, "did": _f.DID})
	if collect.ID.Hex() != mc.Empty {
		result_json := c_code.V1GinError(103, "已收藏，勿重复点击")
		c.JSON(200, result_json)
		return
	}
	//入库
	insert := model.MemberCollect{
		ID:  primitive.NewObjectID(),
		MID: mid,
		DID: _f.DID,
	}
	err := mc.Table(collect.Table()).Insert(insert)
	if err != nil {
		result_json := c_code.V1GinError(104, "添加失败")
		c.JSON(200, result_json)
		return
	}
	model.Movement(mid, index.MID).AddCollect(index)
	result_json := c_code.V1GinSuccess("添加收藏成功")
	c.JSON(200, result_json)
	return
}
