package model

import (
	"fmt"
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"v2ex/app/api"
)

//数据索引
type DataArticle struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Content string             `json:"content" bson:"content"`
	Imgs    []string           `json:"imgs" bson:"imgs"`
	//发布时间
	ReleaseTime time.Time `json:"release_time" bson:"release_time"`
	//调整时间
	ModifyTime    time.Time `json:"modify_time" bson:"modify_time"`
	LastReplyTime time.Time `json:"last_reply_time" bson:"last_reply_time"`
	LastReplyMID  int       `json:"last_reply_mid" bson:"last_reply_mid"`
	CommentSum    int       `json:"comment_sum" bson:"comment_sum"`
	CommentRoot   int       `json:"comment_root" bson:"comment_root"`
	//相关搜索
	RelatedTime time.Time            `json:"related_time" bson:"related_time"`
	RelatedList []primitive.ObjectID `json:"related_list" bson:"related_list"`
}

func (t DataArticle) Table() string {
	return "data_article"
}

//添加文章
func (t DataArticle) Add(mid MIDTYPE, title, content string, insert_time time.Time) gin.H {

	//得到DID
	did, err := AutoID{}.DataID()
	if err != nil && did == 0 {
		return c_code.V1GinError(102, "ID 生成失败")
	}

	//分离图片
	html, imgs, err := api.SeparatePicture(content)
	if err != nil {
		return c_code.V1GinError(102, "处理html错误")
	}
	//定义索引数据
	index := DataIndex{
		ID:    primitive.NewObjectID(),
		DID:   DIDTYPE(did),
		DTYPE: DTYPEArticle,
		MID:   mid,
		T:     title,
		RC:    0,
		CT:    insert_time.Unix(),
	}

	err = mc.Table(index.Table()).Insert(index)
	if err != nil {
		return c_code.V1GinError(103, "写入索引表失败")
	}

	//定义Article 文章表数据
	d_article := DataArticle{
		ID:            index.ID,
		Content:       html,
		Imgs:          imgs,
		ReleaseTime:   insert_time,
		ModifyTime:    insert_time,
		LastReplyTime: time.Time{},
		LastReplyMID:  0,
		CommentSum:    0,
		CommentRoot:   0,
		RelatedTime:   time.Time{},
		RelatedList:   nil,
	}

	err = mc.Table(d_article.Table()).Insert(d_article)
	if err != nil {
		return c_code.V1GinError(104, "写入文件表失败")
	}
	result_json := c_code.V1GinSuccess("", "添加成功", fmt.Sprintf("/a/%d", did))
	AutoID{}.DataAdd()

	//添加进通知中心
	Movement(mid, 0).AddArticleSend(index)
	return result_json
}
