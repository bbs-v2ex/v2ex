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
)

//数据审核类型
const (
	//添加文章
	DataCheckTypeAddArticle = 1
)
