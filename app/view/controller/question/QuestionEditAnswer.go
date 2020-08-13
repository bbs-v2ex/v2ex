package question

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"v2ex/app/view"
)

func QuestionEditAnswer(c *gin.Context) {
	_ht := defaultData(c)
	did, _ := strconv.Atoi(c.Param("did"))
	_ht["index"] = gin.H{
		"DID": did,
	}
	view.Render(c, "question/edit_answer", _ht)
}
