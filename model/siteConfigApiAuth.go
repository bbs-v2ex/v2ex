package model

import "github.com/123456/c_code/mc"
import "github.com/globalsign/mgo/bson"

const _site_api_auth = 2

type SiteConfigApiAuth struct {
	//是否开放注册
	Register bool `json:"register" bson:"register"`
	//爬虫签名
	SpiderSign string `json:"spider_sign" bson:"spider_sign"`

	//文章审核
	ArticleCheck check `json:"article_check" bson:"article_check"`

	//问题审核
	QuestionCheck check `json:"question_check" bson:"question_check"`
}

type check struct {
	//发布
	Send bool `json:"send" bson:"send"`
	//编辑
	Edit bool `json:"edit" bson:"edit"`

	//评论root
	CRootAdd  bool `json:"c_root_add" bson:"c_root_add"`
	CRootEdit bool `json:"c_root_edit" bson:"c_root_edit"`
	//评论child
	CChildAdd  bool `json:"c_child_add" bson:"c_child_add"`
	CChildEdit bool `json:"c_child_edit" bson:"c_child_edit"`
}

func (t SiteConfig) GetApiAuth() (sc SiteConfigApiAuth) {
	mc.Table(t.Table()).Where(bson.M{"key": _site_api_auth}).FindOne(&sc)
	return
}

func (t SiteConfig) SetApiAuth(sc SiteConfigApiAuth) error {
	err := mc.Table(t.Table()).Where(bson.M{"key": _site_api_auth}).UpdateOneIsEmptyNewInsert(&sc)
	return err
}

func (t SiteConfigApiAuth) WaitCheck(user Member, cint int) bool {

	//如果不是注册会员或者是超级管理员则不需要审核直接写入数据库
	if user.MemberType == MemberTypeRoot || !user.IsUser {
		return false
	}

	switch cint {
	case DataCheckTypeArticleAdd:
		return t.ArticleCheck.Send
	case DataCheckTypeArticleEdit:
		return t.ArticleCheck.Edit
	case DataCheckTypeArticleCommentRootAdd:
		return t.ArticleCheck.CRootAdd
	case DataCheckTypeArticleCommentRootEdit:
		return t.ArticleCheck.CRootEdit
	case DataCheckTypeArticleCommentChildAdd:
		return t.ArticleCheck.CChildAdd
	case DataCheckTypeArticleCommentChildEdit:
		return t.ArticleCheck.CChildEdit
	}
	return false
}
