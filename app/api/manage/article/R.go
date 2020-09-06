package article

import (
	"github.com/gin-gonic/gin"
)

func R(r *gin.RouterGroup) {
	r1 := r.Group("/article")
	r1.POST("/add", add)
	r1.POST("/list", list)
	r1.POST("/comment_root_add", comment_root_add)
	r1.POST("/comment_root_list", comment_root_list)
	r1.POST("/comment_child_add", comment_child_add)
	r1.POST("/comment_child_list", comment_child_list)
	r1.POST("/zan_add", zan_add)
	r1.POST("/zan_del", zan_del)
}
