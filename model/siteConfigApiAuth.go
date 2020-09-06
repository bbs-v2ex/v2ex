package model

import "github.com/123456/c_code/mc"
import "github.com/globalsign/mgo/bson"

const _site_api_auth = 2

type SiteConfigApiAuth struct {
	//是否开放注册
	Register bool `json:"register" bson:"register"`
	//爬虫签名
	SpiderSign string `json:"spider_sign" bson:"spider_sign"`
	//发布文章
	SendArticle bool `json:"send_article" bson:"send_article"`
	//文章评论Root
	ArticleCommentRoot bool `json:"article_comment_root" bson:"article_comment_root"`
	//文章评论Child
	ArticleCommentChild bool `json:"article_comment_child" bson:"article_comment_child"`
	//发布问题
	SendQuestion bool `json:"send_question" bson:"send_question"`
	//问题评论Root
	QuestionCommentRoot bool `json:"question_comment_root" bson:"question_comment_root"`
	//问题评论Child
	QuestionCommentChild bool `json:"question_comment_child" bson:"question_comment_child"`
}

func (t SiteConfig) GetApiAuth() (sc SiteConfigApiAuth) {
	mc.Table(t.Table()).Where(bson.M{"key": _site_api_auth}).FindOne(&sc)
	return
}

func (t SiteConfig) SetApiAuth(sc SiteConfigApiAuth) error {
	err := mc.Table(t.Table()).Where(bson.M{"key": _site_api_auth}).UpdateOneIsEmptyNewInsert(&sc)
	return err
}
