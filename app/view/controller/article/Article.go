package article

import (
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"strings"
	api_article "v2ex/app/api/manage/article"
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

	_rid := c.Param("rid")
	rid, err := primitive.ObjectIDFromHex(_rid)

	_ht["comment"] = api_article.CommentRootList(model.DIDTYPE(did), rid, true)
	view.Render(c, "data/article", _ht)
}
