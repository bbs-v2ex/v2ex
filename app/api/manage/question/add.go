package question

import (
	"github.com/123456/c_code"
	"github.com/gin-gonic/gin"
	"time"
	"v2ex/app/api"
	"v2ex/app/nc"
	"v2ex/model"
)

type _add struct {
	Title   string `json:"title"   validate:"required"`
	Content string `json:"content"`
	Html    string `json:"-"`
}

func add(c *gin.Context) {
	_f := _add{}
	c.BindJSON(&_f)
	_f.Title = api.FilterTitle(_f.Title)

	_f.Content = api.FilterContent(_f.Content)

	_f.Html = _f.Content
	_f.Content = c_code.RemoveHtmlTag(_f.Content)
	validator := api.VerifyValidator(_f)
	if validator != "" {
		result_json := c_code.V1GinError(101, validator)
		c.JSON(200, result_json)
		return
	}
	user := api.GetNowUserInfo(c)
	api_auth := model.SiteConfig{}.GetApiAuth()
	//判断是否要审核
	if api_auth.WaitCheck(user, model.DataCheckTypeQuestionAdd) {
		_check := model.DataCheck{

			Type: model.DataCheckTypeQuestionAdd,
			MID:  user.MID,
			DID:  0,
			D: gin.H{
				"title":   _f.Title,
				"content": _f.Html,
			},
		}
		result := model.AddDataCheck(_check)
		c.JSON(200, result)
		return
	}

	err, q_index := nc.QuestionAdd(_f.Title, _f.Html, user.MID, time.Now(), true)
	if err != nil {
		c.JSON(200, c_code.V1GinError(101, err.Error()))
		return
	}
	c.JSON(200, c_code.V1GinSuccess("", "", model.UrlQuestion(q_index)))
}
