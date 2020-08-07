package m_member

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view"
)

func SendArticle(c *gin.Context) {
	_ht := defaultData(c)
	view.Render(c, "m_member/send_article", _ht)
}
