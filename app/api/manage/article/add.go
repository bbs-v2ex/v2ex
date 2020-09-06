package article

import (
	"fmt"
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"v2ex/app/api"
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

	//得到DID
	did, err := model.AutoID{}.DataID()
	if err != nil && did == 0 {
		result_json := c_code.V1GinError(102, "ID 生成失败")
		c.JSON(200, result_json)
		return
	}
	user := api.GetNowUserInfo(c)
	//如果需要审核则进入审核
	if api_auth.SendArticle {
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
		result_json := c_code.V1GinSuccess(200, "已进入后台审核,通过后会展示")
		c.JSON(200, result_json)
		return
	}

	//分离图片
	html, imgs, err := api.SeparatePicture(_f.Html)
	_f.Html = html
	if err != nil {
		result_json := c_code.V1GinError(102, "处理html错误")
		c.JSON(200, result_json)
		return
	}

	//定义索引数据
	index := model.DataIndex{
		ID:    primitive.NewObjectID(),
		DID:   model.DIDTYPE(did),
		DTYPE: model.DTYPEArticle,
		MID:   user.MID,
		T:     _f.Title,
		RC:    0,
		CT:    time.Now().Unix(),
	}

	//定义Article 文章表数据
	d_article := model.DataArticle{
		ID:            index.ID,
		Content:       _f.Html,
		Imgs:          imgs,
		ReleaseTime:   time.Now(),
		ModifyTime:    time.Now(),
		LastReplyTime: time.Time{},
		LastReplyMID:  0,
		CommentSum:    0,
		CommentRoot:   0,
		RelatedTime:   time.Time{},
		RelatedList:   nil,
	}

	err = mc.Table(d_article.Table()).Insert(d_article)
	if err != nil {
		result_json := c_code.V1GinError(104, "写入文件表失败")
		c.JSON(200, result_json)
		return
	}
	result_json := c_code.V1GinSuccess("", "添加成功", fmt.Sprintf("/a/%d", did))
	model.AutoID{}.DataAdd()

	//添加进通知中心
	model.Movement(user.MID, 0).AddArticleSend(index)

	c.JSON(200, result_json)
	return
}
