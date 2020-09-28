package nc

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

//添加文章
func AddArticle(mid model.MIDTYPE, title, content string, insert_time time.Time, find_title_uniqure bool) gin.H {

	if find_title_uniqure {
		//检测title 是否重复
		_title_uniqure := model.DataIndex{}
		mc.Table(_title_uniqure.Table()).Where(bson.M{"t": title}).FindOne(&_title_uniqure)
		if _title_uniqure.ID.Hex() != mc.Empty {
			result_json := c_code.V1GinError(102, "标题重复")
			return result_json
		}
	}

	//得到DID
	did, err := model.AutoID{}.DataID()
	if err != nil && did == 0 {
		return c_code.V1GinError(102, "ID 生成失败")
	}

	//分离图片
	html, imgs, err := api.SeparatePicture(content)
	if err != nil {
		return c_code.V1GinError(102, "处理html错误")
	}
	//定义索引数据
	index := model.DataIndex{
		ID:    primitive.NewObjectID(),
		DID:   model.DIDTYPE(did),
		DTYPE: model.DTYPEArticle,
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
	d_article := model.DataArticle{
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
	model.AutoID{}.DataAdd()

	//添加进通知中心
	model.Movement(mid, 0).AddArticleSend(index)
	return result_json
}
