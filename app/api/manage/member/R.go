package member

import (
	"github.com/gin-gonic/gin"
)

func R(r *gin.RouterGroup) {
	r1 := r.Group("/member")
	r1.POST("/get_user_info", get_user_info)
	r1.POST("/set_user_info", set_user_info)
	r1.POST("/collect_add", collect_add)
	r1.POST("/collect_del", collect_del)
	r1.POST("/is_collect", is_collect)
	//内容审核
	r1.POST("/data_check", data_check)
	r1.POST("/data_check_view", data_check_view)
}
