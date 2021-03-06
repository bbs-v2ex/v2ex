package router

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/api/manage"
	"v2ex/app/api/manage/article"
	"v2ex/app/api/manage/edit"
	"v2ex/app/api/manage/member"
	"v2ex/app/api/manage/question"
	"v2ex/app/api/root_api"
)

func RegisterRoute(r *gin.Engine) {

	//authMiddleware := init_jwt()

	ro := r.Group("/api/manage")

	ro.POST("/add_member", manage.AddMember)
	ro.POST("/add_member_no_user", manage.AddMemberNoUser)
	ro.POST("/login", manage.Login)
	ro.POST("/loginout", manage.LoginOut)
	ro.POST("/show", manage.Show)
	ro.Use(api_send())
	ro.Use(checkLogin)
	ro.POST("/reload_token", manage.ReloadToken)
	root_api.R(ro)
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

	//注册会员
	member.R(ro)

	//注册edit
	edit.R(ro)
}
