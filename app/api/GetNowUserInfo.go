package api

import (
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/model"
)

func GetNowUserInfo(c *gin.Context) (member model.Member) {
	_mid, exists := c.Get("mid")
	if !exists {
		return
	}
	mc.Table(member.Table()).Where(bson.M{"mid": _mid}).FindOne(&member)
	return
}
