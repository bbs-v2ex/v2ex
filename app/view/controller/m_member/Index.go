package m_member

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view"
)

func Index(c *gin.Context) {
	//
	//u := "/_/member/z"
	_ht := defaultData(c)

	_type := c.Param("_type")

	switch _type {
	case "article", "question":
		_ht["index_type"] = _type
		_ht["it"] = _type[:1]
		view.Render(c, "m_member/_index_list", _ht)
		break
	default:
		//u2 := strings.ReplaceAll(u, "/z", "/c")
		_ht["send_list"] = []gin.H{}

		view.Render(c, "m_member/_index", _ht)
		break
	}
}
