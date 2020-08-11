package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/app/api/member"
	"v2ex/app/view"
)

func LastActivity(c *gin.Context) {
	_ht := defaultData(c)
	rid, _ := primitive.ObjectIDFromHex(c.Param("rid"))

	list := member.ListDynamic(0, rid)
	_ht["dt"] = list
	_ht["next_link"] = ""
	if len(list) >= 10 {
		_ht["next_link"] = "/last_activity/" + list[len(list)-1].ID
	}
	if len(list) <= 1 {
		c.Redirect(301, "/last_activity")
	}
	view.Render(c, "_last_activity", _ht)
}
