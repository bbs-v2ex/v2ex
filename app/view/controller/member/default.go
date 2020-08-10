package member

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view/controller"
)

func defaultData(c *gin.Context) gin.H {
	return controller.DefaultData(c)
}
