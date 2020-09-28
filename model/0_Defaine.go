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
	DataCheckTypeAddArticle = 1
)

//会员类型 如果为负 则为异常状态 并限制操作
const (
	//超级管理员
	MemberTypeRoot    = 1
	MemberTypeAverage = 0
)
