package m_config

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view"
)

func Seo(c *gin.Context) {
	_ht := defaultData(c)
	view.Render(c, "m_config/seo", _ht)
}
