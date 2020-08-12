package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"v2ex/app/view/controller"
	"v2ex/app/view/controller/article"
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
	r.GET("/",
		setParam("home"), controller.Home,
	)
	r.GET(
		"/l/:rid",
		setParam("home"),
		controller.Home,
	)
	//动态页
	r.GET(
		"/last_activity",
		setParam("last_activity"),
		controller.LastActivity,
	)
	r.GET("/last_activity/:rid",
		setParam("last_activity"), controller.LastActivity)

	//注册
	r.GET(
		"/registered",
		setParam("registered"), controller.Registered)

	//登录
	r.GET(
		"/login",
		setParam("login"), controller.Login)

	//会员
	_manage := r.Group("/_", setParam("manage"))
	member_center := _manage.Group("/member")

	//会员首页
	member_center.GET("/z/", member_manage.Index)
	member_center.GET("/z/:_type/:_cz", member_manage.Index)
	member_center.GET("/z/:_type", member_manage.Index)
	member_center.GET("/c/:_type", member_manage.CRouter)

	//超级管理员权限页面
	r_config := _manage.Group("/config")
	r_config.GET("/seo", m_config.Seo)

	//文章页
	v_article := r.Group(fmt.Sprintf("/%s", model.UrlTagArticle))
	v_article.GET("/",
		setParam("article_list"), article.Index)

	v_article.GET("/:did/l/:rid",
		setParam("article_list"), article.Index)

	v_article.GET("/:did", setParam("article"),
		article.Article, setParam("article"))

	v_article.GET(fmt.Sprintf("/:did/%s/:rid", model.UrlTagArticleReply),
		setParam("article"), article.Article)

	//问题页
	v_question := r.Group(fmt.Sprintf("/%s", model.UrlTagQuestion))
	v_question.GET("/", question.Index,
		setParam("question_list"))

	v_question.GET("/:did/l/:rid",
		setParam("question_list"), question.Index)

	v_question.GET(fmt.Sprintf("/:did"),
		setParam("question"), question.Question)

	v_question.GET(fmt.Sprintf("/:did/edit_answer"),
		setParam("question_edit"), question.QuestionEditAnswer)

	v_question.GET(fmt.Sprintf("/:did/%s/:rid", model.UrlTagQuestionReply),
		setParam("question"), question.Question)

	//会员页面
	v_member := r.Group(fmt.Sprintf("/%s", model.UrlTagMember))

	v_member.GET("/",
		setParam("member_list"))

	v_member.GET("/:mid",
		setParam("member"), member.Member)

	v_member.GET("/:mid/:_type",
		setParam("member"), member.Member)

}
