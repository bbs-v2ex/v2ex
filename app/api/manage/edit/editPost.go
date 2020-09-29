package edit

import (
	"errors"
	"fmt"
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"v2ex/app/api"
	"v2ex/app/nc"

	"v2ex/model"
)

//判断是否数据是否允许修改允许修改
func editAllow(where bson.M, mid model.MIDTYPE) (err error) {
	//先判断数据库数据是否存在
	data_index := model.DataIndex{}
	mc.Table(data_index.Table()).Where(where).FindOne(&data_index)
	if data_index.MID != mid {
		err = errors.New("请勿乱传参")
		return
	}
	return nil
}

func editPost(c *gin.Context) {
	_f := _edit_result{}
	err := c.Bind(&_f)
	if err != nil {
		result_json := c_code.V1GinError(100, err.Error())
		c.JSON(200, result_json)
		return
	}
	mid := api.GetMID(c)
	user := api.GetNowUserInfo(c)
	api_auth := model.SiteConfig{}.GetApiAuth()
	switch _f.Type {
	//修改文章
	case "article":
		err := editAllow(bson.M{"did": _f.DID, "d_type": model.DTYPEArticle, "mid": mid}, user.MID)
		if err != nil {
			result_json := c_code.V1GinError(10000, err.Error())
			c.JSON(200, result_json)
			return
		}
		//判断是否进入审核阶段
		if api_auth.WaitCheck(user, model.DataCheckTypeArticleEdit) {

			//检测是否已存在提交但是未审核的
			cwhere := bson.M{"type": model.DataCheckTypeArticleEdit, "mid": user.MID, "did": _f.DID}

			_check := model.DataCheck{}
			mc.Table(_check.Table()).Where(cwhere).FindOne(&_check)
			if _check.ID.Hex() != mc.Empty {
				result_json := c_code.V1GinError(20000, "切勿重复提交数据")
				c.JSON(200, result_json)
				return
			}

			edit_article := model.DataCheck{
				ID:   primitive.NewObjectID(),
				Type: model.DataCheckTypeArticleEdit,
				MID:  user.MID,
				DID:  _f.DID,
				D: gin.H{
					"title":   _f.Title,
					"content": _f.Content,
				},
				Itime: time.Now(),
			}

			result := model.AddDataCheck(edit_article)

			c.JSON(200, result)
			return
		} else {
			//不需要审核 直接修改数据库
			err := nc.ArticleEdit(_f.Title, _f.Content, _f.DID, user.MID, time.Now(), true)
			if err != nil {
				result_json := c_code.V1GinError(108, err.Error())
				c.JSON(200, result_json)
				return
			}
			result_json := c_code.V1GinSuccess("修改成功", "", fmt.Sprintf("/%s/%d", model.UrlTagArticle, _f.DID))
			c.JSON(200, result_json)
			return
		}
	case "question":
		//_f.Content = c_code.RemoveHtmlTag(_f.Content)
		//先判断是否可编辑
		err := editAllow(bson.M{"did": _f.DID, "d_type": model.DTYPEQuestion, "mid": mid}, user.MID)
		if err != nil {
			result_json := c_code.V1GinError(10000, err.Error())
			c.JSON(200, result_json)
			return
		}

		//判断是否需要审核
		if api_auth.WaitCheck(user, model.DataCheckTypeQuestionEdit) {
			//检测是否已存在提交但是未审核的
			cwhere := bson.M{"type": model.DataCheckTypeQuestionEdit, "mid": user.MID, "did": _f.DID}

			_check := model.DataCheck{}
			mc.Table(_check.Table()).Where(cwhere).FindOne(&_check)
			if _check.ID.Hex() != mc.Empty {
				result_json := c_code.V1GinError(20000, "切勿重复提交数据")
				c.JSON(200, result_json)
				return
			}
			edit := model.DataCheck{
				ID:   primitive.NewObjectID(),
				Type: model.DataCheckTypeQuestionEdit,
				MID:  user.MID,
				DID:  _f.DID,
				D: gin.H{
					"title":   _f.Title,
					"content": _f.Content,
				},
				Itime: time.Now(),
			}
			result := model.AddDataCheck(edit)
			c.JSON(200, result)
			return
		}
		err = nc.QuestionEdit(_f.Title, _f.Content, _f.DID, user.MID, time.Now(), true)
		if err != nil {
			c.JSON(200, c_code.V1GinError(10000, err.Error()))
			return
		}
		c.JSON(200, c_code.V1GinSuccess("ok"))
		return

	case "question_answer":
		//_f.Content = c_code.RemoveHtmlTag(_f.Content)
		//查询是否存在
		answer := model.CommentQuestionRoot{}
		err := mc.Table(answer.Table()).Where(bson.M{"mid": mid, "did": _f.DID}).FindOne(&answer)

		if err != nil {
			c.JSON(200, c_code.V1GinError(103, "查询失败"))
			return
		}
		if answer.ID.Hex() == mc.Empty {
			c.JSON(200, c_code.V1GinError(104, "查询失败"))
			return
		}

		//检测是否需要审核
		if api_auth.WaitCheck(user, model.DataCheckTypeQuestionCommentRootEdit) {
			//检测是否已存在提交但是未审核的
			cwhere := bson.M{"type": model.DataCheckTypeQuestionCommentRootEdit, "mid": user.MID, "did": _f.DID}

			_check := model.DataCheck{}
			mc.Table(_check.Table()).Where(cwhere).FindOne(&_check)
			if _check.ID.Hex() != mc.Empty {
				result_json := c_code.V1GinError(20000, "切勿重复提交数据")
				c.JSON(200, result_json)
				return
			}

			_check = model.DataCheck{
				Type: model.DataCheckTypeQuestionCommentRootEdit,
				MID:  user.MID,
				DID:  _f.DID,
				D: gin.H{
					"txt": _f.Content,
				},
			}
			result := model.AddDataCheck(_check)
			c.JSON(200, result)

			return
		}

		index := model.DataIndex{}
		mc.Table(index.Table()).Where(bson.M{"d_type": model.DTYPEQuestion, "did": answer.DID}).FindOne(&index)

		err = nc.QuestionCommentRootEdit(_f.Content, _f.DID, user.MID, answer)
		if err != nil {
			c.JSON(200, c_code.V1GinError(105, err.Error()))
			return
		}
		u := model.UrlQuestionAnswer(index, answer)
		c.JSON(200, c_code.V1GinSuccess("修改成功", "", u))
		return
	}

}
