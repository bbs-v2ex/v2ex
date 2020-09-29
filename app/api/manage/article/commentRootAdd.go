package article

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
	"time"
	"v2ex/app/api"
	"v2ex/app/nc"
	"v2ex/model"
)

type _comment_root_add struct {
	DID model.DIDTYPE `json:"did" validate:"gt=0" comment:"文章ID"`
	Txt string        `json:"txt" validate:"min=10,max=1000" comment:"数据"`
}

func commentRootAdd(c *gin.Context) {
	//获取用户信息
	user := api.GetNowUserInfo(c)
	if user.MID == 0 {
		result_json := c_code.V1GinError(101, "非法操作")
		c.JSON(200, result_json)
		return
	}
	_f := _comment_root_add{}
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
	if api_auth.WaitCheck(user, model.DataCheckTypeArticleCommentRootAdd) {

		index := model.DataIndex{}
		mc.Table(index.Table()).Where(bson.M{"did": _f.DID, "d_type": model.DTYPEArticle}).FindOne(&index)
		if index.DID == 0 {
			result_json := c_code.V1GinError(101, "数据丢失")
			c.JSON(200, result_json)
			return
		}

		//需要经过审核
		_check := model.DataCheck{
			Type:  model.DataCheckTypeArticleCommentRootAdd,
			MID:   user.MID,
			DID:   index.DID,
			D:     gin.H{"txt":_f.Txt},
		}
		result := model.AddDataCheck(_check)
		result["u"] = ""
		c.JSON(200,result)
		return
	} else {
		comment_root_id := primitive.NewObjectID()

		err := nc.ArticleCommentRootAdd(_f.DID, _f.Txt, user.MID, time.Now(), comment_root_id)
		if err != nil {
			result_json := c_code.V1GinError(103, err.Error())
			c.JSON(200, result_json)
		}

		ref := c.GetHeader("Referer")
		_u := regexp.MustCompile(`/r/[\w|\s]{24}`).ReplaceAllString(ref, "")
		_u += "/r/" + comment_root_id.Hex()
		result_json := c_code.V1GinSuccess("", "添加成功", _u)

		c.JSON(200, result_json)
	}
	return
}
