package router

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/api/manage"
	"v2ex/app/api/manage/article"
	"v2ex/app/api/manage/question"
	"v2ex/app/api/site_config"
)

func RegisterRoute(r *gin.Engine) {

	//authMiddleware := init_jwt()

	ro := r.Group("/api/manage")

	ro.POST("/add_member", manage.AddMember)
	ro.POST("/login", manage.Login)
	ro.POST("/loginout", manage.LoginOut)

	ro.Use(checkLogin)
	site_config.R(ro)
	//获取用户信息
	ro.POST("/get_user_info", manage.GetUserInfo)

	//临时下载图片
	ro.POST("/download_temp_img", manage.DownloadTempImg)

	//SEO菜单
	ro.POST("/member_nav", manage.MemberNav)

	//注册文章
	article.R(ro)
	//注册提问
	question.R(ro)

}
