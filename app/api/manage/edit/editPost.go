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
				ID:    primitive.NewObjectID(),
				Type:  model.DataCheckTypeArticleEdit,
				MID:   user.MID,
				DID:   _f.DID,
				D:     gin.H{"did": _f.DID, "title": _f.Title, "content": _f.Content},
				Itime: time.Now(),
			}
			//添加进审核表
			err := mc.Table(edit_article.Table()).Insert(edit_article)
			if err != nil {
				result_json := c_code.V1GinError(400, "添加审核表失败")
				c.JSON(200, result_json)
				return
			}
			result_json := c_code.V1GinSuccess(200, "已进入后台审核,通过后会展示", model.UrlViewMemberConfig+"/data_check_view?id="+edit_article.ID.Hex())
			c.JSON(200, result_json)
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
		data_index := model.DataIndex{}
		mc.Table(data_index.Table()).Where(bson.M{"did": _f.DID, "d_type": model.DTYPEQuestion, "mid": mid}).FindOne(&data_index)
		if data_index.MID != mid {
			result_json := c_code.V1GinError(101, "请勿乱传参")
			c.JSON(200, result_json)
			return
		}
		mc.Table(data_index.Table()).Where(bson.M{"_id": data_index.ID}).UpdateOne(bson.M{"t": _f.Title})
		content, imgs, err := api.SeparatePicture(_f.Content)
		if err != nil {
			result_json := c_code.V1GinError(102, "请勿乱传参")
			c.JSON(200, result_json)
			return
		}
		mc.Table(data_index.InfoQuestion.Table()).Where(bson.M{"_id": data_index.ID}).UpdateOne(bson.M{"content": content, "imgs": imgs})

		result_json := c_code.V1GinSuccess("修改成功", "", fmt.Sprintf("/%s/%d", model.UrlTagQuestion, _f.DID))
		c.JSON(200, result_json)
		return

	case "question_answer":
		_html, _imgs, err2 := api.SeparatePicture(_f.Content)
		if err2 != nil {
			result_json := c_code.V1GinError(102, "html解析错误")
			c.JSON(200, result_json)
			return
		}

		mid := api.GetMID(c)
		answer := model.CommentQuestionRoot{}
		err := mc.Table(answer.Table()).Where(bson.M{"mid": mid, "did": _f.DID}).FindOne(&answer)
		if err != nil || answer.ID.Hex() == mc.Empty {
			result_json := c_code.V1GinError(103, "查询失败")
			c.JSON(200, result_json)
			return
		}
		update := bson.M{
			"text": _html,
			"img":  _imgs,
		}
		err = mc.Table(answer.Text.Table()).Where(bson.M{"_id": answer.ID}).UpdateOne(update)
		if err != nil {
			result_json := c_code.V1GinError(104, "修改失败")
			c.JSON(200, result_json)
			return
		}
		result_json := c_code.V1GinSuccess("修改成功", "", fmt.Sprintf("/%s/%d/%s/%s", model.UrlTagQuestion, _f.DID, model.UrlTagQuestionReply, answer.ID.Hex()))
		c.JSON(200, result_json)
		return
	}

}
