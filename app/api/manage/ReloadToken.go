package manage

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/model"
	"v2ex/until"
)

func ReloadToken(c *gin.Context) {
	token := c.GetHeader("____token")
	err := mc.Table(model.MemberToken{}.Table()).Where(bson.M{"token": token}).UpdateOne(bson.M{"expire": until.MemberTokenAddValidPeriod()})
	if err != nil {
		result_json := c_code.V1GinError(101, "")
		c.JSON(200, result_json)
		return
	}
	result_json := c_code.V1GinSuccess("")
	c.JSON(200, result_json)
}
