package root_api

import (
	"github.com/123456/c_code"
	"github.com/gin-gonic/gin"
)

func dataSearchNav(c *gin.Context) {

	result := gin.H{
		"s_list": []gin.H{
			{
				"t": "全部",
				"v": "",
			},
			{
				"t": "文章",
				"v": "article",
			},
			{
				"t": "问题",
				"v": "question",
			},
		},
		"search": gin.H{
			"mid":   0,
			"title": "",
			"type":  "",
		},
	}

	c.JSON(200, c_code.V1GinSuccess(result))
}
