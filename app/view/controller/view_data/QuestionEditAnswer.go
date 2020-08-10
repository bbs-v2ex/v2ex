package view_data

import (
	"github.com/gin-gonic/gin"
	"v2ex/app/view"
)

func QuestionEditAnswer(c *gin.Context) {
	_ht := defaultData(c)
	view.Render(c, "data/question_edit_answer", _ht)
}
