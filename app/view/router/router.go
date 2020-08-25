package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"v2ex/app/view"
	"v2ex/app/view/controller"
	"v2ex/app/view/controller/article"
	"v2ex/app/view/controller/edit"
	"v2ex/app/view/controller/m_config"
	member_manage "v2ex/app/view/controller/m_member"
	"v2ex/app/view/controller/member"
	"v2ex/app/view/controller/question"
	"v2ex/model"
)

func setParam(view_type string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("view_type", view_type)
	}
}

func RegisterRoute(r *gin.Engine) {

	//robots 文件
	r.StaticFile("/robots.txt", "./000_robots.txt")

	r.GET("/jump-address",
		setParam(view.ViewTypeJumtAadress), controller.JumtAddress)

	r.GET("/",
		setParam(view.ViewTypeHome), controller.Home,
	)
	r.GET(
		"/l/:rid",
		setParam(view.ViewTypeHome),
		controller.Home,
	)
	//动态页
	r.GET(
		"/last_activity",
		setParam(view.ViewTypeLastActivity),
		controller.LastActivity,
	)
	r.GET("/last_activity/:rid",
		setParam(view.ViewTypeLastActivity), controller.LastActivity)

	//注册
	r.GET(
		"/registered",
		setParam(view.ViewTypeRegistered), controller.Registered)

	//登录
	r.GET(
		"/login",
		setParam(view.ViewTypeLogin), controller.Login)

	//会员
	_manage := r.Group("/_", setParam(view.ViewTypeManage))
	member_center := _manage.Group("/member")

	//会员首页
	member_center.GET("/z/", member_manage.Index)
	member_center.GET("/z/:_type/:_cz", member_manage.Index)
	member_center.GET("/z/:_type", member_manage.Index)
	member_center.GET("/c/:_type", member_manage.CRouter)

	//超级管理员权限页面
	r_config := _manage.Group("/config")
	r_config.GET("/*seo", m_config.Index)

	//文章页
	v_article := r.Group(fmt.Sprintf("/%s", model.UrlTagArticle))
	v_article.GET("/",
		setParam(view.ViewTypeArticleList), article.Index)

	v_article.GET("/:did/l/:rid",
		setParam(view.ViewTypeArticleList), article.Index)

	v_article.GET("/:did", setParam("article"),
		article.Article, setParam(view.ViewTypeArticle))

	v_article.GET(fmt.Sprintf("/:did/%s/:rid", model.UrlTagArticleReply),
		setParam(view.ViewTypeArticle), article.Article)

	//问题页
	v_question := r.Group(fmt.Sprintf("/%s", model.UrlTagQuestion))
	v_question.GET("/",
		setParam(view.ViewTypeQuestionList), question.Index)

	v_question.GET("/:did/l/:rid",
		setParam(view.ViewTypeQuestionList), question.Index)

	v_question.GET(fmt.Sprintf("/:did"),
		setParam(view.ViewTypeQuestion), question.Question)

	v_question.GET(fmt.Sprintf("/:did/%s/:rid", model.UrlTagQuestionReply),
		setParam(view.ViewTypeQuestion), question.Question)

	//会员页面
	v_member := r.Group(fmt.Sprintf("/%s", model.UrlTagMember))

	v_member.GET("/",
		setParam(view.ViewTypeMemberList), member.Index)

	v_member.GET("/:mid",
		setParam(view.ViewTypeMember), member.Member)

	v_member.GET("/:mid/:_type",
		setParam(view.ViewTypeMember), member.Member)
	//更新网站地图接口
	r.GET("/update_site_map", view.UpdateSiteMap)
	//sitemap 文件位置
	r.Static("/site_map_check", "./__sitemap")
	//最新的 1000 条数据
	r.GET("/last_data_1000", view.LastData1000)
	r.GET("/login_auto", controller.LoginAuto)
	r.GET("/edit/edit", edit.Edit)
}
