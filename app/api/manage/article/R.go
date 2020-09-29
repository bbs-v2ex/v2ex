package article

import (
	"github.com/gin-gonic/gin"
)

func R(r *gin.RouterGroup) {
	r1 := r.Group("/article")
	r1.POST("/add", add)
	r1.POST("/list", list)
	r1.POST("/comment_root_add", commentRootAdd)
	r1.POST("/comment_root_list", commentRootList)
	r1.POST("/comment_child_add", commentChildAdd)
	r1.POST("/comment_child_list", commentChildList)
	r1.POST("/zan_add", zanAdd)
	r1.POST("/zan_del", zanDel)
}
