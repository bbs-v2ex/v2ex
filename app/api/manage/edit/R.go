package edit

import (
	"github.com/gin-gonic/gin"
)

func R(r *gin.RouterGroup) {
	r1 := r.Group("/edit")
	r1.GET("/edit", edit)
	r1.POST("/edit", editpost)
}
