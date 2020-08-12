package member

import (
	"fmt"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"path/filepath"
	"strconv"
	"v2ex/app/api/member"
	"v2ex/app/view"
	"v2ex/model"
)

func Member(c *gin.Context) {
	_ht := defaultData(c)
	mid, _ := strconv.Atoi(c.Param("mid"))

	u := fmt.Sprintf("/member/%d", mid)
	_member_mav := []gin.H{
		{
			"t":      "动态",
			"u":      u,
			"active": false,
		},
		{
			"t":      "提问",
			"u":      u + "/question",
			"active": false,
		},
		{
			"t":      "回答",
			"u":      u + "/comment",
			"active": false,
		},
		{
			"t":      "文章",
			"u":      u + "/article",
			"active": false,
		},
		{
			"t":      "收藏",
			"u":      u + "/collect",
			"active": false,
		},
	}
	_type := c.Param("_type")
	_type_active := false
	for k, _ := range _member_mav {
		_f_last := filepath.Base(_member_mav[k]["u"].(string))
		if _f_last == _type {
			_type_active = true
			_member_mav[k]["active"] = true
			break
		}
	}

	if !_type_active {
		_member_mav[0]["active"] = true
	}
	_ht["_member_mav"] = _member_mav

	//查询是否存在此会员
	user_info := model.Member{}.GetUserInfo(model.MIDTYPE(mid), true)
	if user_info.MID == 0 {
		view.R404(c, view.ViewError{Message: "无此会员"})
		return
	}
	_ht["user_info"] = user_info

	//页面分发
	tpl_name := ""
	txt := ""
	switch _type {
	case "question":
		txt = "进行提问"
		_ht["dt"] = member.ListQuestion(model.MIDTYPE(mid), primitive.ObjectID{})
		tpl_name = "member/question"
		break
	case "comment":
		txt = "对问题进行回复"
		_ht["dt"] = member.ListComment(model.MIDTYPE(mid), primitive.ObjectID{})
		tpl_name = "member/user_home"
		break
	case "article":
		txt = "发布过文章"
		_ht["dt"] = member.ListArticle(model.MIDTYPE(mid), primitive.ObjectID{})
		tpl_name = "member/question"
		break
	case "collect":
		txt = "收藏"
		_ht["dt"] = member.ListCollect(model.MIDTYPE(mid), primitive.ObjectID{})
		tpl_name = "member/collect"
		break
	default:
		_ht["member_body_empty"] = "还没有动态哦"
		_ht["dt"] = member.ListDynamic(model.MIDTYPE(mid), primitive.ObjectID{})
		tpl_name = "member/user_home"
		break
	}

	//获取最新文章
	index := model.DataIndex{}
	_article_list := []model.DataIndex{}
	mc.Table(index.Table()).Where(bson.M{"d_type": model.DTYPEArticle}).Limit(10).Order(bson.M{"_id": -1}).Find(&_article_list)
	article_list := []gin.H{}
	for _, v := range _article_list {
		article_list = append(article_list, gin.H{
			"t": v.T,
			"u": model.UrlArticle(v),
		})
	}
	_ht["article_list"] = article_list

	//获取最新的提问
	_question_list := []model.DataIndex{}
	mc.Table(index.Table()).Where(bson.M{"d_type": model.DTYPEQuestion}).Limit(10).Order(bson.M{"_id": -1}).Find(&_question_list)
	question_list := []gin.H{}
	for _, v := range _question_list {
		question_list = append(question_list, gin.H{
			"t": v.T,
			"u": model.UrlQuestion(v),
		})
	}
	_ht["question_list"] = question_list

	_ht["member_body_empty"] = "还没有" + txt + "哦"
	view.Render(c, tpl_name, _ht)
}
