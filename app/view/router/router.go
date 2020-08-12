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

func RegisterRoute(r *gin.Engine) {
	r.GET("/", controller.Home)

	//动态页
	r.GET("/last_activity", controller.LastActivity)
	r.GET("/last_activity/:rid", controller.LastActivity)

	//注册
	r.GET("/registered", controller.Registered)

	//登录
	r.GET("/login", controller.Login)

	//会员
	member_center := r.Group("/_/member")

	//会员首页
	member_center.GET("/z/", member_manage.Index)
	member_center.GET("/z/:_type/:_cz", member_manage.Index)
	member_center.GET("/z/:_type", member_manage.Index)

	//文章发布
	member_center.GET("/c/:_type", member_manage.CRouter)

	//超级管理员权限页面
	r_config := r.Group("/_/config")
	r_config.GET("/seo", m_config.Seo)

	//文章页
	v_article := r.Group(fmt.Sprintf("/%s", model.UrlTagArticle))
	v_article.GET("/", article.Index)
	v_article.GET("/:did", article.Article)
	v_article.GET(fmt.Sprintf("/:did/%s/:rid", model.UrlTagArticleReply), article.Article)

	//问题页
	v_question := r.Group(fmt.Sprintf("/%s", model.UrlTagQuestion))
	v_question.GET("/", question.Index)
	v_question.GET(fmt.Sprintf("/:did"), question.Question)
	v_question.GET(fmt.Sprintf("/:did/edit_answer"), question.QuestionEditAnswer)
	v_question.GET(fmt.Sprintf("/:did/%s/:rid", model.UrlTagQuestionReply), question.Question)

	//会员页面
	v_member := r.Group(fmt.Sprintf("/%s", model.UrlTagMember))
	v_member.GET("/")
	v_member.GET("/:mid", member.Index)
	v_member.GET("/:mid/:_type", member.Index)

}
