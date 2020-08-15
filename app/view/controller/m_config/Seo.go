package m_config

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view"
)

func Index(c *gin.Context) {
	_ht := defaultData(c)
	file_tpl := "m_config" + c.Param("seo")
	view.Render(c, file_tpl, _ht)
}
