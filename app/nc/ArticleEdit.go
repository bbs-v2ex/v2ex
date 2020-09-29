package nc

import (
	"errors"
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	"time"
	"v2ex/app/api"
	"v2ex/model"
)

func ArticleEdit(title, content string, did model.DIDTYPE, mid model.MIDTYPE, update_time time.Time, find_title_uniqure bool) error {

	data_index := model.DataIndex{}
	mc.Table(data_index.Table()).Where(bson.M{"did": did, "d_type": model.DTYPEArticle, "mid": mid}).FindOne(&data_index)
	if data_index.ID.Hex() == mc.Empty {
		return errors.New("未找到数据")
	}

	if find_title_uniqure {
		//检测title 是否重复
		_title_uniqure := model.DataIndex{}
		mc.Table(_title_uniqure.Table()).Where(bson.M{"did": bson.M{"$ne": did}, "t": title, "d_type": model.DTYPEArticle}).FindOne(&_title_uniqure)
		if _title_uniqure.ID.Hex() != mc.Empty {
			return errors.New("标题重复")
		}
	}
	//分离图片
	html, imgs, err := api.SeparatePicture(content)
	if err != nil {
		return errors.New("处理html错误")
	}
	//先改索引表的标题
	err = mc.Table(model.DataIndex{}.Table()).Where(bson.M{"_id": data_index.ID}).UpdateOne(bson.M{"t": title})
	if err != nil {
		return err
	}
	//在修改文章表
	err = mc.Table(model.DataIndex{}.InfoArticle.Table()).Where(bson.M{"_id": data_index.ID}).UpdateOne(bson.M{"content": html, "imgs": imgs})
	if err != nil {
		return err
	}
	return nil
}
