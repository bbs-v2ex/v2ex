package manage

import (
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/model"
)

type _show struct {
	DID model.DIDTYPE `json:"did" validate:"gt=0" comment:"文章ID"`
}

func Show(c *gin.Context) {
	_f := _show{}
	c.BindJSON(&_f)
	if _f.DID == 0 {
		return
	}
	mc.Table(model.DataIndex{}.Table()).Where(bson.M{"did": _f.DID}).FieldAddOrDel("show", +1)
}
