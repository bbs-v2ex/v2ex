package router

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/app/nc"
	"v2ex/model"
)

func api_send() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 随机获取一个会员并设置为 mid
		spider_sign := c.GetHeader("spider_sign")
		if spider_sign != "" {
			api_auth := nc.GetApiAuth()
			if api_auth.SpiderSign == spider_sign {
				list_mid := []model.Member{}
				mc.Table(model.Member{}.Table()).Where(bson.M{"is_user": false}).Projection(bson.M{"mid": 1}).Find(&list_mid)
				if len(list_mid) > 0 {
					c.Set("mid", list_mid[c_code.Rand(len(list_mid)-1)].MID)
				} else {
					c.Set("mid", model.MIDTYPE(27))
				}

			}
		}
	}
}
