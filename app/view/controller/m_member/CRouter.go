package m_member

import (
	"github.com/gin-gonic/gin"
	"strings"
	"v2ex/app/view"
)

func CRouter(c *gin.Context) {
	_ht := defaultData(c)
	_type := strings.TrimSpace(c.Param("_type"))

	switch _type {
	case "article", "question":
		_ht["index_type"] = _type
		_ht["it"] = _type[:1]
		view.Render(c, "m_member/_index_list", _ht)
		return
	case "/", "":
		view.Render(c, "m_member/_index_list", _ht)
		return
	default:
		//u2 := strings.ReplaceAll(u, "/z", "/c")
		//_ht["send_list"] = []gin.H{}
		//
		//view.Render(c, "m_member/_index", _ht)
		break
	}

	view.Render(c, "m_member/"+_type, _ht)
}
