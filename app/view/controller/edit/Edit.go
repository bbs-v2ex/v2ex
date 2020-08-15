package edit

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view"
	"v2ex/model"
)

type _edit struct {
	DID  model.DIDTYPE `form:"did" json:"did"`
	Type string        `form:"type" json:"type"`
}

func Edit(c *gin.Context) {
	_f := _edit{}
	c.Bind(&_f)
	_ht := defaultData(c)
	_ht["f1"] = _f
	view.Render(c, "edit/"+_f.Type, _ht)
}
