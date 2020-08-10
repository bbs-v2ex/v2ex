package manage

import (
	"fmt"
	"github.com/123456/c_code"
	"github.com/gin-gonic/gin"
	"v2ex/app/api"
)

func MemberNav(c *gin.Context) {
	user_info := api.GetNowUserInfo(c)
	u2 := "/_/member/c"
	list := []gin.H{
		{
			"t": "我的主页",
			"u": fmt.Sprintf("/member/%d", user_info.MID),
		},
		{
			"t": "发文章",
			"u": u2 + "/send_article",
		},

		{
			"t": "发提问",
			"u": u2 + "/send_question",
		},
		{
			"t": "个人资料",
			"u": u2 + "/user_info",
		},
	}
	if user_info.MemberType == 1 {
		u := "/_/config/"
		list = append(list, []gin.H{
			{
				"t": "SEO设置",
				"u": u + "seo",
			},
		}...)

	} else {

	}
	//登录成功
	result_json := c_code.V1GinSuccess(list)
	c.JSON(200, result_json)
	return
}
