package question

import (
	"github.com/123456/c_code"
	"github.com/gin-gonic/gin"
	"v2ex/app/api"
)

func R(r *gin.RouterGroup) {
	r1 := r.Group("/question")
	r1.POST("/add", add)
	r1.POST("/list", list)
	r1.POST("/comment_root_add", comment_root_add)
	r1.POST("/is_root_edit", is_root_edit)
	r1.POST("/get_self_answer", get_self_answer)
	r1.POST("/edit_self_answer", edit_self_answer)
	//r1.POST("/comment_root_list", comment_root_list)
	//r1.POST("/comment_child_add", comment_child_add)
	//r1.POST("/comment_child_list", comment_child_list)
	//r1.POST("/zan_add", zan_add)
	//r1.POST("/zan_del", zan_del)
}

func isok(c *gin.Context) {
	user := api.GetNowUserInfo(c)
	if user.MemberType != 1 {
		result := c_code.V1GinError(500, "没权限啊")
		c.JSON(200, result)
		c.Abort()
		return
	}
}
