package m_member

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
	"v2ex/app/view"
)

func Index(c *gin.Context) {

	u := "/_/member"
	_ht := defaultData(c)
	_member_mav := []gin.H{
		{
			"t":      "默认选中",
			"u":      u + "/index",
			"active": false,
		},
		{
			"t":      "我的文章",
			"u":      u + "/article",
			"active": false,
		},
		{
			"t":      "我的提问",
			"u":      u + "/question",
			"active": false,
		},
	}

	_type := c.Param("_type")
	_type_active := false
	for k, _ := range _member_mav {
		_f_last := filepath.Base(_member_mav[k]["u"].(string))
		if _f_last == _type {
			_type_active = true
			_member_mav[k]["active"] = true
			break
		}
	}

	if !_type_active {
		_member_mav[0]["active"] = true
	}
	_ht["_member_mav"] = _member_mav
	switch _type {
	case "article", "question":
		view.Render(c, "m_member/_index_list", _ht)
		break
	default:

		_ht["send_list"] = []gin.H{
			{
				"t": "发文章",
				"u": u + "/send_article",
			},
			{
				"t": "发提问",
				"u": u + "/send_question",
			},
		}

		view.Render(c, "m_member/_index", _ht)
		break
	}
}
