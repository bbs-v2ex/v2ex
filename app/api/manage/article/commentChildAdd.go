package article

import (
	"github.com/123456/c_code"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"v2ex/app/api"
	"v2ex/app/nc"
	"v2ex/model"
)

type _comment_child_add struct {
	RID primitive.ObjectID `json:"rid"  bson:"rid" validate:"len=12" comment:"评论ID"`
	PID primitive.ObjectID `json:"pid" bson:"pid"`
	Txt string             `json:"txt" bson:"txt"  validate:"min=10,max=1000" comment:"数据"`
}

func commentChildAdd(c *gin.Context) {
	//获取用户信息
	user := api.GetNowUserInfo(c)
	if user.MID == 0 {
		result_json := c_code.V1GinError(101, "非法操作")
		c.JSON(200, result_json)
		return
	}
	_f := _comment_child_add{}
	c.BindJSON(&_f)
	validator := api.VerifyValidator(_f)
	if validator != "" {
		result_json := c_code.V1GinError(102, validator)
		c.JSON(200, result_json)
		return
	}
	//删除html标签
	_f.Txt = c_code.RemoveHtmlTag(_f.Txt)

	api_auth := model.SiteConfig{}.GetApiAuth()

	//检测是否需要经过审核
	if api_auth.WaitCheck(user, model.DataCheckTypeArticleCommentChildAdd) {

		_check := model.DataCheck{
			Type: model.DataCheckTypeArticleCommentChildAdd,
			MID:  user.MID,
			DID:  0,
			D: gin.H{
				"rid": _f.RID,
				"pid": _f.PID,
				"txt": _f.Txt,
			},
		}
		result := model.AddDataCheck(_check)
		result["u"] = ""
		c.JSON(200, result)
		return
	}
	err := nc.ArticleCommentChildAdd(_f.RID, _f.PID, _f.Txt, user.MID, time.Now(), primitive.NewObjectID())
	if err != nil {
		result_json := c_code.V1GinError(103, err.Error())
		c.JSON(200, result_json)
	}
	result_json := c_code.V1GinSuccess("", "添加成功")

	c.JSON(200, result_json)

}
