package model

import (
	"errors"
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	bson2 "go.mongodb.org/mongo-driver/bson"
	"v2ex/app/common"
)

type MovementHtml struct {
	// 最上方得提示信息
	ST string
	//作者信息
	Author struct {
		Name   string
		Avatar string
		Des    string
	}
	//多少人赞同
	Zan   int
	TextS struct {
		H     string
		Imags []string
	}
	Text string
}

func (t MovementCenter) ToConversion() (hs MovementHtml, err error) {
	var json []byte
	json, err = bson.MarshalJSON(t.V)
	if err != nil {
		return
	}
	switch t.Type {
	case MovementArticleSend: //1
		d := articleSend{}
		err = bson2.UnmarshalExtJSON(json, false, &d)
		if err != nil {
			return
		}

		//直接组装html
		index := DataIndex{}
		mc.Table(index.Table()).Where(bson.M{"did": d.DID}).FindOne(&index)
		if index.DID == 0 {
			err = errors.New("文章丢失")
			return
		}
		hs.ST = "发布了文章"
		//获取数据详情
		err = mc.Table(index.InfoArticle.Table()).Where(bson.M{"_id": index.ID}).FindOne(&index.InfoArticle)
		if index.InfoArticle.ID.Hex() == mc.Empty {
			err = errors.New("文章丢失")
			return
		}
		hs.TextS.H = index.InfoArticle.Content
		hs.TextS.Imags = index.InfoArticle.Imgs
		break

	case MovementArticleCommentRoot: // 2
		d := articleCommentRoot{}
		err = bson2.UnmarshalExtJSON(json, false, &d)
		if err != nil {
			return
		}
		//获取数据
		comment_article_root := CommentRoot{}
		err = mc.Table(comment_article_root.Table()).Where(bson.M{"_id": d.RID}).FindOne(&comment_article_root)
		if err != nil || comment_article_root.ID.Hex() == mc.Empty {
			err = errors.New("回复丢失")
			return
		}
		//获取文本
		err = mc.Table(comment_article_root.Text.Table()).Where(bson.M{"_id": comment_article_root.ID}).FindOne(&comment_article_root.Text)
		if err != nil || comment_article_root.Text.ID.Hex() == mc.Empty {
			err = errors.New("回复丢失")
			return
		}

		hs.ST = "对文章发布评论"

		//获取作者信息
		if t.MID != comment_article_root.MID {
			author := Member{}.GetUserInfo(comment_article_root.MID, true)
			hs.Author = struct {
				Name   string
				Avatar string
				Des    string
			}{Name: author.UserName, Avatar: common.Avatar(author.Avatar), Des: author.More.Des}
		}
		//多少人赞同
		hs.Zan = comment_article_root.ZanLen
		//封装文本
		hs.TextS = struct {
			H     string
			Imags []string
		}{H: comment_article_root.Text.Text, Imags: comment_article_root.Text.Img}

		break

	case MovementArticleCommentGood: //3
		d := articleCommentZan{}
		err = bson2.UnmarshalExtJSON(json, false, &d)
		if err != nil {
			return
		}
		comment_article_root := CommentRoot{}
		err = mc.Table(comment_article_root.Table()).Where(bson.M{"_id": d.RID}).FindOne(&comment_article_root)
		if err != nil || comment_article_root.ID.Hex() == mc.Empty {
			err = errors.New("回复丢失")
			return
		}
		//获取文本
		err = mc.Table(comment_article_root.Text.Table()).Where(bson.M{"_id": comment_article_root.ID}).FindOne(&comment_article_root.Text)
		if err != nil || comment_article_root.Text.ID.Hex() == mc.Empty {
			err = errors.New("回复丢失")
			return
		}
		hs.ST = "对文章评论赞同"
		//获取作者信息
		if t.MID != comment_article_root.MID {
			author := Member{}.GetUserInfo(comment_article_root.MID, true)
			hs.Author = struct {
				Name   string
				Avatar string
				Des    string
			}{Name: author.UserName, Avatar: common.Avatar(author.Avatar), Des: author.More.Des}
		}
		//多少人赞同
		hs.Zan = comment_article_root.ZanLen
		//封装文本
		hs.TextS = struct {
			H     string
			Imags []string
		}{H: comment_article_root.Text.Text, Imags: comment_article_root.Text.Img}
		break

	case MovementQuestionSend: //4
		d := questionSend{}
		err = bson2.UnmarshalExtJSON(json, false, &d)
		if err != nil {
			return
		}
		//直接组装html
		index := DataIndex{}
		mc.Table(index.Table()).Where(bson.M{"did": d.DID}).FindOne(&index)
		if index.DID == 0 {
			err = errors.New("提问丢失")
			return
		}
		hs.ST = "发布了问题"
		//获取数据详情
		mc.Table(index.InfoQuestion.Table()).Where(bson.M{"_id": index.ID}).FindOne(&index.InfoQuestion)
		if index.InfoQuestion.ID.Hex() == mc.Empty {
			err = errors.New("提问丢失")
			return
		}
		hs.TextS.H = index.InfoQuestion.Content
		hs.TextS.Imags = index.InfoQuestion.Imgs
		break

	case MovementQuestionCommentRoot: //5
		d := questionCommentRoot{}
		err = bson2.UnmarshalExtJSON(json, false, &d)
		if err != nil {
			return
		}
		//获取数据
		comment_question_root := CommentQuestionRoot{}
		err = mc.Table(comment_question_root.Table()).Where(bson.M{"_id": d.RID}).FindOne(&comment_question_root)
		if err != nil || comment_question_root.ID.Hex() == mc.Empty {
			err = errors.New("we丢失")
			return
		}
		//获取文本
		err = mc.Table(comment_question_root.Text.Table()).Where(bson.M{"_id": comment_question_root.ID}).FindOne(&comment_question_root.Text)
		if err != nil || comment_question_root.Text.ID.Hex() == mc.Empty {
			err = errors.New("回复丢失")
			return
		}

		hs.ST = "对问题进行回复"

		//获取作者信息
		if t.MID != comment_question_root.MID {
			author := Member{}.GetUserInfo(comment_question_root.MID, true)
			hs.Author = struct {
				Name   string
				Avatar string
				Des    string
			}{Name: author.UserName, Avatar: common.Avatar(author.Avatar), Des: author.More.Des}
		}
		//多少人赞同
		hs.Zan = comment_question_root.ZanLen
		//封装文本
		hs.TextS = struct {
			H     string
			Imags []string
		}{H: comment_question_root.Text.Text, Imags: comment_question_root.Text.Img}

		break

	case MovementQuestionCommentGood: //3
		d := questionAnswerZan{}
		err = bson2.UnmarshalExtJSON(json, false, &d)
		if err != nil {
			return
		}
		comment_question_root := CommentQuestionRoot{}
		err = mc.Table(comment_question_root.Table()).Where(bson.M{"_id": d.RID}).FindOne(&comment_question_root)
		if err != nil || comment_question_root.ID.Hex() == mc.Empty {
			err = errors.New("回复丢失")
			return
		}
		//获取文本
		err = mc.Table(comment_question_root.Text.Table()).Where(bson.M{"_id": comment_question_root.ID}).FindOne(&comment_question_root.Text)
		if err != nil || comment_question_root.Text.ID.Hex() == mc.Empty {
			err = errors.New("回复丢失")
			return
		}
		hs.ST = "对回答赞同"
		//获取作者信息
		if t.MID != comment_question_root.MID {
			author := Member{}.GetUserInfo(comment_question_root.MID, true)
			hs.Author = struct {
				Name   string
				Avatar string
				Des    string
			}{Name: author.UserName, Avatar: common.Avatar(author.Avatar), Des: author.More.Des}
		}
		//多少人赞同
		hs.Zan = comment_question_root.ZanLen
		//封装文本
		hs.TextS = struct {
			H     string
			Imags []string
		}{H: comment_question_root.Text.Text, Imags: comment_question_root.Text.Img}
		break
	}
	return
}
