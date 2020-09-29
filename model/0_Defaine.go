package model

type MIDTYPE int64
type DIDTYPE int64

const DTYPEArticle = 1
const DTYPEQuestion = 2

//Url 链接 状态
const (
	UrlTagArticle       = "a"
	UrlTagArticleReply  = "r"
	UrlTagQuestion      = "q"
	UrlTagQuestionReply = "answer"
	UrlTagMember        = "member"

	//Url 普通管理界面的Url前缀
	UrlViewMemberConfig = "/_/member/c"

	//Url 普通管理界面的 管理 Url前缀
	UrlViewMemberManage = "/_/config"
)

//数据审核类型
const (
	//添加文章
	DataCheckTypeArticleAdd = 1
	//编辑文章
	DataCheckTypeArticleEdit = 2
	//评论Root Add
	DataCheckTypeArticleCommentRootAdd = 3
	//评论Root Add
	DataCheckTypeArticleCommentRootEdit  = 4
	DataCheckTypeArticleCommentChildAdd  = 5
	DataCheckTypeArticleCommentChildEdit = 6

	//问题类
	//添加提问
	DataCheckTypeQuestionAdd              = 100
	DataCheckTypeQuestionEdit             = 101
	DataCheckTypeQuestionCommentRootAdd   = 103
	DataCheckTypeQuestionCommentRootEdit  = 104
	DataCheckTypeQuestionCommentChildAdd  = 105
	DataCheckTypeQuestionCommentChildEdit = 106
)

//会员类型 如果为负 则为异常状态 并限制操作
const (
	//超级管理员
	MemberTypeRoot    = 1
	MemberTypeAverage = 0
)
