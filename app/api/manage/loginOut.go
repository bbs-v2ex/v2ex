package manage

import (
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/model"
)

func LoginOut(c *gin.Context) {
	token := c.GetHeader("____token")
	member_token := model.MemberToken{}
	mc.Table(member_token.Table()).Where(bson.M{"token": token}).DelOne()
}
