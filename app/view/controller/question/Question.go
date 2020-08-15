package question

import (
	"fmt"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"strings"
	"time"
	"v2ex/app/api"
	api_question "v2ex/app/api/manage/question"
	"v2ex/app/view"
	"v2ex/app/view/controller"
	"v2ex/model"
	"v2ex/until"
)

func Question(c *gin.Context) {

	did, _ := strconv.Atoi(c.Param("did"))
	t_list := []string{}
	q_where := bson.M{"did": did}

	qid_len := len(c.Param("did"))
	switch {
	case qid_len == 24:
		hex, err := primitive.ObjectIDFromHex(c.Param("did"))
		if err != nil {
			view.R404(c, view.ViewError{Message: "问题不存在"})
			return
		}
		q_where = bson.M{"_id": hex}
		break
	case did == 0:
		view.R404(c, view.ViewError{Message: "问题不存在"})
		return
	}
	q_where["d_type"] = model.DTYPEQuestion
	//查询数据库
	question := model.DataQuestion{}
	index := model.DataIndex{}
	err := mc.Table(index.Table()).Where(q_where).FindOne(&index)
	if err != nil {
		view.R404(c, view.ViewError{Message: "问题不存在111"})
		return
	}
	if index.DID == 0 {
		view.R404(c, view.ViewError{Message: "问题不存在222"})
		return
	}
	//查询文章详细数据
	mc.Table(question.Table()).Where(bson.M{"_id": index.ID}).FindOne(&question)
	if question.ID.Hex() == mc.Empty {
		view.R404(c, view.ViewError{Message: "文章不存在333"})
		return
	}
	index.InfoQuestion = question
	index.InfoQuestion.Content = api.RestorePicture(index.InfoQuestion.Content, index.T, index.InfoQuestion.Imgs, "600")
	//渲染数据
	_ht := defaultData(c)
	_ht["index"] = index
	mt := model.Member{}
	member_info := mt.GetUserInfo(index.MID, true)
	_ht["member_info"] = member_info

	_rid := c.Param("rid")
	rid, err := primitive.ObjectIDFromHex(_rid)

	_ht["comment"] = api_question.CommentRootList(model.DIDTYPE(index.DID), rid, true)

	t_list = append(t_list, index.T)
	_ht["t"] = controller.TitleJoin(t_list)
	k := controller.KeywordJoin(index.T + index.InfoQuestion.Content)
	_ht["k"] = k

	d := model.DesSplit(index.InfoQuestion.Content, 120)
	if len(d) < 30 {
		d = index.T
	}
	_ht["d"] = d
	_ht["sp_t"] = index.T

	//获取右边的数据

	//加载相关文章

	nids := []primitive.ObjectID{
		index.ID,
	}
	//获取最新文章
	_vd_new := []model.DataIndex{}
	vd_new := []gin.H{}
	mc.Table(index.Table()).Where(bson.M{"d_type": model.DTYPEQuestion, "_id": bson.M{"$nin": nids}}).Order(bson.M{"_id": -1}).Limit(10).Find(&_vd_new)
	for _, v := range _vd_new {
		vd_new = append(vd_new, gin.H{
			"t": v.T,
			"u": model.UrlQuestion(v),
		})
		nids = append(nids, v.ID)
	}
	_ht["vd_new"] = vd_new
	//检测是否需要更新相关文章列表
	r_list := []gin.H{}
	_r_list := []model.DataIndex{}
	if index.InfoQuestion.RelatedTime.Unix() < until.DataTimeDifference(-7).Unix() {
		//if true {
		//需要更新相关数列表
		r_ci_list := strings.Split(k, "，")
		if len(r_ci_list) == 0 {
			r_ci_list = []string{"的", "问"}
		}
		r_where := bson.M{
			"_id": bson.M{"$nin": nids},
			"t":   bson.M{"$regex": strings.Join(r_ci_list, "|")},
		}
		mc.Table(index.Table()).Where(r_where).Order(bson.M{"_id": -1}).Limit(10).Find(&_r_list)
		rl_list := []primitive.ObjectID{}
		fmt.Println(r_where)
		for _, v := range _r_list {
			rl_list = append(rl_list, v.ID)
		}
		mc.Table(model.DataQuestion{}.Table()).Where(bson.M{"_id": index.ID}).UpdateOne(bson.M{"related_time": time.Now(), "related_list": rl_list})

	} else {
		if len(index.InfoQuestion.RelatedList) >= 1 {
			mc.Table(index.Table()).Where(bson.M{"_id": bson.M{"$in": index.InfoQuestion.RelatedList}}).Find(&_r_list)
		}
	}

	for _, v := range _r_list {

		_one := gin.H{}
		_one["t"] = v.T
		if v.DTYPE == model.DTYPEQuestion {
			_one["u"] = model.UrlQuestion(v)
		}
		if v.DTYPE == model.DTYPEArticle {
			_one["u"] = model.UrlArticle(v)
		}
		r_list = append(r_list, _one)
		nids = append(nids, v.ID)
	}
	_ht["vd_rl"] = r_list

	view.Render(c, "question/data", _ht)
}
