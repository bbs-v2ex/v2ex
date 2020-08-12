package question

import (
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"strings"
	api_question "v2ex/app/api/manage/question"
	"v2ex/app/view"
	"v2ex/model"
)

func Question(c *gin.Context) {

	_type := c.Query("type")

	switch _type {
	case "edit_answer":
		c.Redirect(302, "/q/"+c.Param("did")+"/edit_answer")
		return
	}

	did, _ := strconv.Atoi(c.Param("did"))
	t_list := []string{}
	if did == 0 {
		view.R404(c, view.ViewError{Message: "问题不存在"})
		return
	}
	//查询数据库
	question := model.DataQuestion{}
	index := model.DataIndex{}
	err := mc.Table(index.Table()).Where(bson.M{"did": did}).FindOne(&index)
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
	//渲染数据
	_ht := defaultData(c)
	_ht["index"] = index
	t_list = append(t_list, index.T)
	t_list = append(t_list, _ht["t_"].(string))
	_ht["t"] = strings.Join(t_list, _ht["title_fgf"].(string))

	mt := model.Member{}
	member_info := mt.GetUserInfo(index.MID, true)
	_ht["member_info"] = member_info

	_rid := c.Param("rid")
	rid, err := primitive.ObjectIDFromHex(_rid)

	_ht["comment"] = api_question.CommentRootList(model.DIDTYPE(did), rid, true)
	view.Render(c, "data/question", _ht)
}
