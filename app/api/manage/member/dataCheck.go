package member

import (
	"fmt"
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/app/api"
	"v2ex/model"
)

func dataCheck(c *gin.Context) {
	mid := api.GetMID(c)
	//查询是否是管理员
	member := model.Member{}.GetUserInfo(mid, true)
	where := bson.M{"mid": mid}
	if member.MemberType == model.MemberTypeRoot {
		where = bson.M{}
	}
	//查询数据
	list := []model.DataCheck{}
	mc.Table(model.DataCheck{}.Table()).Where(where).Limit(30).Find(&list)
	l := []gin.H{}
	for _, v := range list {
		_one := gin.H{
			"time": c_code.StrTime(v.Itime),
		}
		_cz := []gin.H{
			{
				"t":    "查看",
				"u":    "data_check_view?id=" + v.ID.Hex(),
				"ajax": false,
			},
		}
		c_title := ""
		switch v.Type {
		case model.DataCheckTypeArticleAdd:
			c_title += "发布文章"
			break
		case model.DataCheckTypeArticleEdit:
			c_title += "修改文章 " + fmt.Sprintf("[%d]", v.DID)
			break
		case model.DataCheckTypeArticleCommentRootAdd:
			c_title += "ArticleCommentRootAdd"
			break
		case model.DataCheckTypeArticleCommentChildAdd:
			c_title += "ArticleCommentChildAdd"
			break
		case model.DataCheckTypeQuestionAdd:
			c_title += "发布提问"
			break
		case model.DataCheckTypeQuestionEdit:
			c_title += "编辑提问"
			break
		case model.DataCheckTypeQuestionCommentRootAdd:
			c_title += "回答回答"
			break
		case model.DataCheckTypeQuestionCommentRootEdit:
			c_title += "编辑回答"
			break
		case model.DataCheckTypeQuestionCommentChildAdd:
			c_title += "QuestionCommentChildAdd"
			break
		case model.DataCheckTypeQuestionCommentChildEdit:
			c_title += "QuestionCommentChildEdit"
			break
		}
		_one["title"] = c_title
		_one["cz"] = _cz
		l = append(l, _one)
	}
	c.JSON(200, c_code.V1GinSuccess(l))
}
