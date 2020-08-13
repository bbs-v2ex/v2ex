package view

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/nc"
	"v2ex/model"
)

const (
	ViewTypeHome         = "home"
	ViewTypeLastActivity = "last_activity"
	ViewTypeRegistered   = "registered"
	ViewTypeLogin        = "login"
	ViewTypeManage       = "manage"
	ViewTypeArticleList  = "article_list"
	ViewTypeArticle      = "article"
	ViewTypeQuestionList = "question_list"
	ViewTypeQuestion     = "question"
	ViewTypeQuestionEdit = "question_edit"
	ViewTypeMemberList   = "member_list"
	ViewTypeMember       = "member"
)

func setNavigation(c *gin.Context, _ht gin.H) (n []gin.H) {

	seoconfig := nc.GetSeoConfig()
	_navigation := []gin.H{
		{
			"t":   seoconfig.NavigationHomeTitle,
			"u":   model.Url("/"),
			"tag": true,
		},
	}
	_t, exists := c.Get("view_type")
	if !exists {
		return
	}
	switch _t.(string) {
	case ViewTypeHome:
		break
	case ViewTypeLastActivity:
		_navigation = append(_navigation, gin.H{
			"t":   "网站动态",
			"u":   model.Url("/last_activity"),
			"tag": false,
		})
		return _navigation
	}
	return
}
