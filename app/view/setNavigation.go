package view

import (
	"fmt"
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
	case ViewTypeRegistered:
		_navigation = append(_navigation, []gin.H{
			{
				"t":   "会员",
				"u":   model.Url("/" + model.UrlTagMember),
				"tag": true,
			},
			{
				"t":   "注册",
				"u":   "",
				"tag": false,
			},
		}...)
		return _navigation
	case ViewTypeLogin:
		_navigation = append(_navigation, []gin.H{
			{
				"t":   "会员",
				"u":   model.Url("/" + model.UrlTagMember),
				"tag": true,
			},
			{
				"t":   "登录",
				"u":   "",
				"tag": false,
			},
		}...)
		return _navigation
	case ViewTypeManage:
		_navigation = append(_navigation, []gin.H{
			{
				"t":   "管理",
				"u":   "/_/member/z/",
				"tag": false,
			},
		}...)
		return _navigation
	case ViewTypeArticleList:
		_navigation = append(_navigation, []gin.H{
			{
				"t":   "文章列表",
				"u":   model.Url(fmt.Sprintf("/%s/", model.UrlTagArticle)),
				"tag": false,
			},
		}...)
		return _navigation
	case ViewTypeQuestionList:
		_navigation = append(_navigation, []gin.H{
			{
				"t":   "问题列表",
				"u":   model.Url(fmt.Sprintf("/%s/", model.UrlTagQuestion)),
				"tag": false,
			},
		}...)
		return _navigation
	case ViewTypeArticle:
		_navigation = append(_navigation, []gin.H{
			{
				"t":   "文章",
				"u":   model.Url(fmt.Sprintf("/%s/", model.UrlTagArticle)),
				"tag": true,
			},
			{
				"t":   _ht["sp_t"],
				"u":   "",
				"tag": false,
			},
		}...)
		return _navigation
	case ViewTypeMember:
		_navigation = append(_navigation, []gin.H{
			{
				"t":   "会员",
				"u":   model.Url(fmt.Sprintf("/%s/", model.UrlTagMember)),
				"tag": true,
			},
			{
				"t":   "",
				"u":   "",
				"tag": false,
			},
		}...)
		return _navigation
	}
	return
}
