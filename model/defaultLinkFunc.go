package model

import (
	"fmt"
	"github.com/123456/c_code"
	"regexp"
	"strings"
	"v2ex/config"
)

const (
	UrlTagArticle       = "a"
	UrlTagArticleReply  = "r"
	UrlTagQuestion      = "q"
	UrlTagQuestionReply = "answer"
	UrlTagMember        = "member"
)

func DesSplit(s string, sp int) string {
	s = c_code.RemoveHtmlTag(s)
	s = regexp.MustCompile(`[\r\n|\n|\r|\t|{xxx{img}xxx}]+`).ReplaceAllString(s, " ")
	runes := []rune(s)
	if len(runes) > sp {
		s = string(runes[:sp]) + "..."
	}
	return s
}

func UrlArticle(index DataIndex) string {
	return fmt.Sprintf("/%s/%d", UrlTagArticle, index.DID)
}
func UrlArticleAnswer(index DataIndex, comment_root CommentRoot) string {
	return fmt.Sprintf("/%s/%d/%s/%s", UrlTagArticle, index.DID, UrlTagArticleReply, comment_root.ID.Hex())
}

func UrlQuestion(index DataIndex) string {
	return fmt.Sprintf("/%s/%d", UrlTagQuestion, index.DID)
}
func UrlQuestionAnswer(index DataIndex, comment_root CommentQuestionRoot) string {
	return fmt.Sprintf("/%s/%d/%s/%s", UrlTagQuestion, index.DID, UrlTagQuestionReply, comment_root.ID.Hex())
}
func UrlMember(index Member) string {
	return fmt.Sprintf("/%s/%d", UrlTagMember, index.MID)
}
func UrlImage(string2 ...string) string {
	u := strings.Join(string2, "/")
	if !strings.HasPrefix(u, "/") {
		return u
	}

	_con := config.GetConfig()

	return fmt.Sprintf("%s%s", _con.Run.UploadServer, u)
}