package question

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
	"v2ex/app/api"
	"v2ex/app/nc"
	"v2ex/model"
)

type _comment_root_add struct {
	DID  model.DIDTYPE `json:"did" validate:"gt=0" comment:"文章ID"`
	Txt  string        `json:"txt"`
	Text string        `validate:"min=10,max=100000000" comment:"数据"`
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
	_f.Text = c_code.RemoveHtmlTag(_f.Txt)
	validator := api.VerifyValidator(_f)
	if validator != "" {
		result_json := c_code.V1GinError(102, validator)
		c.JSON(200, result_json)
		return
	}
	//验证是否回答过此问题
	is_reply := model.CommentQuestionRoot{}
	err := mc.Table(is_reply.Table()).Where(bson.M{"mid": user.MID, "did": _f.DID}).FindOne(&is_reply)
	if err != nil {
		result_json := c_code.V1GinError(103, "查询失败")
		c.JSON(200, result_json)
		return
	}

	if is_reply.ID.Hex() != mc.Empty {
		result_json := c_code.V1GinError(104, "不允许重复回答")
		c.JSON(200, result_json)
		return
	}

	//检测是否需要审核
	api_auth := model.SiteConfig{}.GetApiAuth()
	if api_auth.WaitCheck(user, model.DataCheckTypeQuestionCommentRootAdd) {
		//检测是否提交过修改请求
		cwhere := bson.M{
			"type": model.DataCheckTypeQuestionCommentRootAdd,
			"mid":  user.MID,
			"did":  _f.DID,
		}
		_check := model.DataCheck{}
		mc.Table(_check.Table()).Where(cwhere).FindOne(&_check)
		if _check.ID.Hex() != mc.Empty {
			result_json := c_code.V1GinError(104, "提交的回答正在审核中,不可以重复提交")
			c.JSON(200, result_json)
			return
		}
		_check = model.DataCheck{
			Type: model.DataCheckTypeQuestionCommentRootAdd,
			MID:  user.MID,
			DID:  _f.DID,
			D: gin.H{
				"txt": _f.Txt,
			},
		}
		result := model.AddDataCheck(_check)
		c.JSON(200, result)
		return
	}

	comment_id := primitive.NewObjectID()
	err = nc.QuestionCommentRootAdd(_f.Txt, _f.DID, user.MID, comment_id)
	if err != nil {
		result_json := c_code.V1GinError(105, err.Error())
		c.JSON(200, result_json)
		return
	}
	ref := c.GetHeader("Referer")
	_u := regexp.MustCompile(`/`+model.UrlTagQuestionReply+`/[\w|\s]{24}`).ReplaceAllString(ref, "")
	_u += "/" + model.UrlTagQuestionReply + "/" + comment_id.Hex()
	result_json := c_code.V1GinSuccess(comment_id, "添加成功", _u)
	c.JSON(200, result_json)
	return
}
