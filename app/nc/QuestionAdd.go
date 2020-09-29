package nc

import (
	"errors"
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"v2ex/app/api"
	"v2ex/model"
)

func QuestionAdd(title, content string, mid model.MIDTYPE, l_time time.Time, find_title_uniqure bool) (err error, index model.DataIndex) {
	if find_title_uniqure {
		//检测title 是否重复
		_title_uniqure := model.DataIndex{}
		mc.Table(_title_uniqure.Table()).Where(bson.M{"t": title}).FindOne(&_title_uniqure)
		if _title_uniqure.ID.Hex() != mc.Empty {
			err = errors.New("标题重复")
			return
		}
	}

	//获取QAID
	did, err := model.AutoID{}.QAID()
	if err != nil && did == 0 {

		return
	}
	//分离图片
	html, imgs, err := api.SeparatePicture(content)
	if err != nil {
		err = errors.New("处理html错误")
		return
	}

	//定义索引数据
	index = model.DataIndex{
		ID:    primitive.NewObjectID(),
		DID:   model.DIDTYPE(did),
		DTYPE: model.DTYPEQuestion,
		MID:   mid,
		T:     title,
		RC:    0,
		CT:    l_time.Unix(),
	}
	err = mc.Table(index.Table()).Insert(index)
	if err != nil {

		err = errors.New("写入索引表失败")
		return
	}

	//定义Article 文章表数据
	d_article := model.DataQuestion{
		ID:            index.ID,
		Content:       html,
		Imgs:          imgs,
		ReleaseTime:   l_time,
		ModifyTime:    l_time,
		LastReplyTime: time.Time{},
		LastReplyMID:  0,
		CommentSum:    0,
		CommentRoot:   0,
		RelatedTime:   time.Time{},
		RelatedList:   nil,
	}
	err = mc.Table(d_article.Table()).Insert(d_article)
	if err != nil {
		err = errors.New("写入文件表失败")
		return
	}
	//result_json := c_code.V1GinSuccess("", "添加成功", fmt.Sprintf("/q/%d", did))
	model.AutoID{}.QAAdd()
	model.Movement(mid, 0).AddQuestionSend(index)
	return
}
