package member

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view"
)

func Index(c *gin.Context) {
	_ht := defaultData(c)
	view.Render(c, "member/user_home", _ht)
}
