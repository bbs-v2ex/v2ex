package article

import (
	"github.com/123456/c_code"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/app/api"
	"v2ex/model"
)

type _comment_root_list struct {
	DID model.DIDTYPE      `json:"did" validate:"gt=0" comment:"文章ID"`
	RID primitive.ObjectID `json:"rid" validate:"len=12" comment:"评论ID"`
}

func comment_root_list(c *gin.Context) {
	_f := _comment_root_list{}
	c.BindJSON(&_f)
	validator := api.VerifyValidator(_f)
	if validator != "" {
		result_json := c_code.V1GinError(101, validator)
		c.JSON(200, result_json)
		return
	}
	comment := CommentRootList(_f.DID, _f.RID)
	result_json := c_code.V1GinSuccess(comment)
	c.JSON(200, result_json)
}
