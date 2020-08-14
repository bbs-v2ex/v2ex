package m_config

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view"
)

func Index(c *gin.Context) {
	_ht := defaultData(c)
	view.Render(c, "m_config/"+c.Param("seo"), _ht)
}
