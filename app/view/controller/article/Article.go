package article

import (
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"strings"
	"time"
	"v2ex/app/api"
	api_article "v2ex/app/api/manage/article"
	"v2ex/app/view"
	"v2ex/app/view/controller"
	"v2ex/model"
	"v2ex/until"
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
	//图片归位
	for k, _ := range index.InfoArticle.Imgs {
		index.InfoArticle.Imgs[k] += "?h=500"
	}
	index.InfoArticle.Content = api.RestorePicture(index.InfoArticle.Content, index.T, index.InfoArticle.Imgs)
	_ht["art"] = index
	t_list = append(t_list, index.T)
	_ht["t"] = controller.TitleJoin(t_list)
	k := controller.KeywordJoin(index.T + index.InfoArticle.Content)
	_ht["k"] = k

	d := model.DesSplit(index.InfoArticle.Content, 120)
	if len(d) < 30 {
		d = index.T
	}
	_ht["d"] = d
	_ht["sp_t"] = index.T
	mt := model.Member{}
	member_info := mt.GetUserInfo(index.MID, true)
	_ht["member_info"] = member_info

	_rid := c.Param("rid")
	rid, err := primitive.ObjectIDFromHex(_rid)

	_ht["comment"] = api_article.CommentRootList(model.DIDTYPE(did), rid, true)
	//加载相关文章

	nids := []primitive.ObjectID{
		index.ID,
	}
	//获取最新文章
	_vd_new := []model.DataIndex{}
	vd_new := []gin.H{}
	mc.Table(index.Table()).Where(bson.M{"d_type": model.DTYPEArticle, "_id": bson.M{"$nin": nids}}).Limit(10).Find(&_vd_new)
	for _, v := range _vd_new {
		vd_new = append(vd_new, gin.H{
			"t": v.T,
			"u": model.UrlArticle(v),
		})
		nids = append(nids, v.ID)
	}
	_ht["vd_new"] = vd_new
	//检测是否需要更新相关文章列表
	r_list := []gin.H{}
	_r_list := []model.DataIndex{}
	if index.InfoArticle.RelatedTime.Unix() < until.DataTimeDifference(-7).Unix() {
		//需要更新相关数列表
		r_ci_list := strings.Split(k, "，")
		if len(r_ci_list) == 0 {
			r_ci_list = []string{"问", "文"}
		}

		mc.Table(index.Table()).Where(bson.M{"_id": bson.M{"$nin": nids}, "t": bson.M{"$regex": "/" + strings.Join(r_ci_list, "|") + "/"}}).Order(bson.M{"_id": -1}).Limit(10).Find(&_r_list)
		rl_list := []model.DIDTYPE{}
		for _, v := range _r_list {
			rl_list = append(rl_list, v.DID)
		}
		mc.Table(model.DataArticle{}.Table()).Where(bson.M{"_id": index.ID}).UpdateOne(bson.M{"related_time": time.Now(), "related_list": rl_list})

	} else {
		if len(index.InfoArticle.RelatedList) >= 1 {
			mc.Table(index.Table()).Where(bson.M{"did": bson.M{"$in": index.InfoArticle.RelatedList}}).Find(&_r_list)
		}
	}

	for _, v := range _r_list {
		r_list = append(r_list, gin.H{
			"t": v.T,
			"u": model.UrlArticle(v),
		})
		nids = append(nids, v.ID)
	}
	_ht["vd_rl"] = r_list
	view.Render(c, "data/article", _ht)
}
