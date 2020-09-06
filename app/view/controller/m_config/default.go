package m_config

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view/controller"
)

func defaultData(c *gin.Context) gin.H {
	data := controller.DefaultData(c)
	data["ID"] = c.Query("id")
	return data
}
