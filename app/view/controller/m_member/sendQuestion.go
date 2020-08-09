package m_member

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view"
)

func SendQuestion(c *gin.Context) {
	_ht := defaultData(c)
	view.Render(c, "m_member/send_question", _ht)
}
