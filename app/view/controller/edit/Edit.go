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
	tpl_name := ""
	switch _f.Type {
	case "question", "article":
		tpl_name = "main"
		break
	case "question_answer":
		tpl_name = "reply"
		break
	}
	view.Render(c, "edit/"+tpl_name, _ht)
}
