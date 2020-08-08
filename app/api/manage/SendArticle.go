package manage

import (
	"fmt"
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"v2ex/app/api"
	"v2ex/model"
)

type _send_article struct {
	Title   string `json:"title"   validate:"required"`
	Content string `json:"content"  validate:"required"`
	Html    string `json:"-"`
}

func SendArticle(c *gin.Context) {
	_f := _send_article{}
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
	//插入数据表
	did, err := model.AutoID{}.DataID()
	if err != nil && did == 0 {
		result_json := c_code.V1GinError(102, "ID 生成失败")
		c.JSON(200, result_json)
		return
	}
	user := api.GetNowUserInfo(c)

	//定义索引数据
	index := model.DataIndex{
		ID:    primitive.NewObjectID(),
		DID:   model.DIDTYPE(did),
		DTYPE: model.DTYPEArticle,
		MID:   user.MID,
		T:     _f.Title,
		RC:    0,
	}
	err = mc.Table(index.Table()).Insert(index)
	if err != nil {
		result_json := c_code.V1GinError(103, "写入索引表失败")
		c.JSON(200, result_json)
		return
	}

	//定义Article 文章表数据
	d_article := model.DataArticle{
		ID:            index.ID,
		Content:       _f.Html,
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
	c.JSON(200, result_json)
	return
}
