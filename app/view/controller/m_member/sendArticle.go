package m_member

import (
	"github.com/gin-gonic/gin"
	"strings"
	"v2ex/app/view"
)

func CRouter(c *gin.Context) {
	_ht := defaultData(c)
	_type := strings.TrimSpace(c.Param("_type"))
	view.Render(c, "m_member/"+_type, _ht)
}
