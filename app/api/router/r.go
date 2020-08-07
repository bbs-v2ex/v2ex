package router

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/api/manage"
)

func RegisterRoute(r *gin.Engine) {

	//authMiddleware := init_jwt()

	ro := r.Group("/api/manage")
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
}
