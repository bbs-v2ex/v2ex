package router

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view/controller"
	member2 "v2ex/app/view/controller/m_member"
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
}
