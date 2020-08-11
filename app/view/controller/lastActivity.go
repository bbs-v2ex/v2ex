package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/app/api/member"
	"v2ex/app/view"
)

func LastActivity(c *gin.Context) {
	_ht := defaultData(c)
	list := member.ListDynamic(0, primitive.ObjectID{})
	_ht["dt"] = list
	view.Render(c, "_last_activity", _ht)
}
