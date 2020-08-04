package manage

import (
	"github.com/123456/c_code"
	"github.com/gin-gonic/gin"
	"v2ex/app/api"
)

//import (
//	"fmt"
//	jwt "github.com/appleboy/gin-jwt/v2"
//	"github.com/gin-gonic/gin"
//)
//var identityKey = "user-mid"
//func GetUserInfo(c *gin.Context)  {
//	claims := jwt.ExtractClaims(c)
//	mid :=  int( claims[identityKey].(float64))
//	fmt.Println(claims,mid)
//}

func GetUserInfo(c *gin.Context) {

	//获取用信息
	now_member := api.GetNowUserInfo(c)
	if now_member.MID == 0 {
		result_json := c_code.V1GinError(600, "MID不存在")
		c.JSON(200, result_json)
		return
	}

	result_data := gin.H{
		"mid":    now_member.MID,
		"name":   now_member.UserName,
		"avatar": now_member.Avatar,
	}
	result_json := c_code.V1GinSuccess(result_data)
	c.JSON(200, result_json)
}
