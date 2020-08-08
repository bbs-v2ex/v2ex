package router

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/api/manage"
	"v2ex/app/api/site_config"
)

func RegisterRoute(r *gin.Engine) {

	//authMiddleware := init_jwt()

	ro := r.Group("/api/manage")
	site_config.R(ro)
	ro.POST("/add_member", manage.AddMember)
	ro.POST("/login", manage.Login)
	ro.POST("/loginout", manage.LoginOut)

	ro.Use(checkLogin)

	//获取用户信息
	ro.POST("/get_user_info", manage.GetUserInfo)
	//发布文章
	ro.POST("/send_article", manage.SendArticle)
	//临时下载图片
	ro.POST("/download_temp_img", manage.DownloadTempImg)

	//SEO菜单
	ro.POST("/member_nav", manage.MemberNav)

}
