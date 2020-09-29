package article

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"v2ex/app/api"
	"v2ex/app/nc"
	"v2ex/model"
)

type _add struct {
	Title   string `json:"title"   validate:"required"`
	Content string `json:"content"  validate:"required"`
	Html    string `json:"-"`
}

func add(c *gin.Context) {

	api_auth := model.SiteConfig{}.GetApiAuth()

	_f := _add{}
	c.BindJSON(&_f)
	_f.Title = api.FilterTitle(_f.Title)

	_f.Content = api.FilterContent(_f.Content)

	_f.Html = _f.Content
	_f.Content = c_code.RemoveHtmlTag(_f.Content)
	validator := api.VerifyValidator(_f)
	if validator != "" {
		result_json := c_code.V1GinError(101, validator)
		c.JSON(200, result_json)
		return
	}

	//检测title 是否重复
	_title_uniqure := model.DataIndex{}
	mc.Table(_title_uniqure.Table()).Where(bson.M{"t": _f.Title}).FindOne(&_title_uniqure)
	if _title_uniqure.ID.Hex() != mc.Empty {
		result_json := c_code.V1GinError(102, "标题重复")
		c.JSON(200, result_json)
		return
	}

	user := api.GetNowUserInfo(c)
	//如果需要审核则进入审核

	if api_auth.WaitCheck(user, model.DataCheckTypeAddArticle) {
		//添加进审核表
		data_check := model.DataCheck{
			ID:    primitive.NewObjectID(),
			Type:  model.DataCheckTypeAddArticle,
			Itime: time.Now(),
			MID:   user.MID,
			D: gin.H{
				"title":   _f.Title,
				"content": _f.Html,
			},
		}
		err := mc.Table(data_check.Table()).Insert(data_check)
		if err != nil {
			result_json := c_code.V1GinError(400, "添加审核表失败")
			c.JSON(200, result_json)
			return
		}
		result_json := c_code.V1GinSuccess(200, "已进入后台审核,通过后会展示", model.UrlViewMemberConfig+"/data_check_view?id="+data_check.ID.Hex())
		c.JSON(200, result_json)
		return
	}
	c.JSON(200, nc.AddArticle(user.MID, _f.Title, _f.Html, time.Now(), false))
	return
}
