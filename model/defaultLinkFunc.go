package model

import (
	"fmt"
)

const (
	UrlTagArticle       = "a"
	UrlTagArticleReply  = "r"
	UrlTagQuestion      = "q"
	UrlTagQuestionReply = "answer"
)

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
