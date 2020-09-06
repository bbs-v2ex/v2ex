package manage

import (
	"fmt"
	"github.com/123456/c_code"
	"github.com/gin-gonic/gin"
	"v2ex/app/api"
	"v2ex/model"
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
		{
			"t": "审核中的内容",
			"u": u2 + "/data_check",
		},
	}
	if user_info.MemberType == model.MemberTypeRoot {
		u := "/_/config/"
		list = append(list, []gin.H{
			{
				"t": "SEO设置",
				"u": u + "seo",
			},
			{
				"t": "网站权限",
				"u": u + "api_auth",
			},
			{
				"t": "索引查看",
				"u": u + "db_index",
			},
		}...)

	} else {

	}
	//登录成功
	result_json := c_code.V1GinSuccess(list)
	c.JSON(200, result_json)
	return
}
