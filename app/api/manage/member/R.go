package member

import (
	"github.com/gin-gonic/gin"
)

func R(r *gin.RouterGroup) {
	r1 := r.Group("/member")
	r1.POST("/get_user_info", getUserInfo)
	r1.POST("/set_user_info", setUserInfo)
	r1.POST("/collect_add", collectAdd)
	r1.POST("/collect_del", collectDel)
	r1.POST("/is_collect", isCollect)
	//内容审核
	r1.POST("/data_check", dataCheck)
	r1.POST("/data_check_view", dataCheckView)
}
