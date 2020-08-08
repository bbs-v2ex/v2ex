package view_data

import (
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"strconv"
	"strings"
	"v2ex/app/view"
	"v2ex/model"
)

func Article(c *gin.Context) {
	did, _ := strconv.Atoi(c.Param("did"))
	t_list := []string{}
	if did == 0 {
		view.R404(c, view.ViewError{Message: "文章不存在"})
		return
	}
	//查询数据库
	article := model.DataArticle{}
	index := model.DataIndex{}
	err := mc.Table(index.Table()).Where(bson.M{"did": did}).FindOne(&index)
	if err != nil {
		view.R404(c, view.ViewError{Message: "文章不存在111"})
		return
	}
	if index.DID == 0 {
		view.R404(c, view.ViewError{Message: "文章不存在222"})
		return
	}
	//查询文章详细数据
	mc.Table(article.Table()).Where(bson.M{"_id": index.ID}).FindOne(&article)
	if article.ID.Hex() == mc.Empty {
		view.R404(c, view.ViewError{Message: "文章不存在333"})
		return
	}
	index.InfoArticle = article
	//渲染数据
	_ht := defaultData(c)
	_ht["art"] = index
	t_list = append(t_list, index.T)
	t_list = append(t_list, _ht["t_"].(string))
	_ht["t"] = strings.Join(t_list, _ht["title_fgf"].(string))

	mt := model.Member{}
	member_info := mt.GetUserInfo(index.MID, true)
	_ht["member_info"] = member_info

	//查询评论
	comment_list := []model.CommentRoot{}
	mc.Table(model.CommentRoot{}.Table()).Where(bson.M{"did": index.DID}).Limit(10).Find(&comment_list)
	_ht["comment"] = comment_list

	view.Render(c, "data/article", _ht)
}
