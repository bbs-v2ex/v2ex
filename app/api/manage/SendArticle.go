package manage

import (
	"github.com/123456/c_code"
	"github.com/gin-gonic/gin"
	"v2ex/app/api"
)

type _send_article struct {
	Title   string `json:"title"   validate:"required"`
	Content string `json:"content"  validate:"required"`
	Html    string `json:"-"`
}

func SendArticle(c *gin.Context) {
	_f := _send_article{}
	c.BindJSON(&_f)
	_f.Html = _f.Content
	_f.Content = c_code.RemoveHtmlTag(_f.Content)
	validator := api.VerifyValidator(_f)
	if validator != "" {
		result_json := c_code.V1GinError(101, validator)
		c.JSON(200, result_json)
		return
	}
	//插入数据表
}
