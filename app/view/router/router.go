package router

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view/controller"
	"v2ex/app/view/controller/m_config"
	member2 "v2ex/app/view/controller/m_member"
	"v2ex/app/view/controller/view_data"
)

func RegisterRoute(r *gin.Engine) {
	r.GET("/", controller.Home)

	//注册
	r.GET("/registered", controller.Registered)

	//登录
	r.GET("/login", controller.Login)

	//会员
	member := r.Group("/_/member")

	//会员首页
	member.GET("/z/", member2.Index)
	member.GET("/z/:_type/:_cz", member2.Index)
	member.GET("/z/:_type", member2.Index)

	//文章发布
	member.GET("/send_article", member2.SendArticle)
	member.GET("/send_question", member2.SendQuestion)

	//超级管理员权限页面
	r_config := r.Group("/_/config")
	r_config.GET("/seo", m_config.Seo)

	//文章页
	r.GET("/a/:did", view_data.Article)
	r.GET("/a/:did/r/:rid", view_data.Article)

	//问题页
	r.GET("/q/:did", view_data.Question)
	r.GET("/q/:did/answer/:rid", view_data.Question)
}
