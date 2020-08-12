package controller

import (
	"github.com/123456/c_code"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/app/api/member"
	"v2ex/app/view"
)

func LastActivity(c *gin.Context) {

	xx := c.Query("xx")

	_ht := defaultData(c)
	rid, _ := primitive.ObjectIDFromHex(c.Param("rid"))

	list := member.ListDynamic(0, rid)
	_ht["dt"] = list
	_ht["next_link"] = ""

	if len(list) >= 10 {
		_ht["next_link"] = "/last_activity/" + list[len(list)-1].ID
	}
	if xx == "nohead" {
		html_content := view.RenderGetContent("member/member_dz_activity.html", _ht)
		_ht["content"] = c_code.CompressHtml(html_content)
		result_json := c_code.V1GinSuccess(html_content)
		result_json["next"] = _ht["next_link"]
		c.JSON(200, result_json)
		return
	}
	if len(list) <= 1 {
		c.Redirect(301, "/last_activity")
	}
	//加载热门问题

	view.Render(c, "_last_activity", _ht)
}
