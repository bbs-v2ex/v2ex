package model

import (
	"errors"
	"fmt"
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
)

func (t MovementCenter) About(mid MIDTYPE) (list []MovementCenter) {
	_list := []MovementCenter{}

	//where := bson.M{
	//	"$or": []bson.M{
	//		{"mid": mid},
	//		//{"m2id": mid},
	//	},
	//}
	where := bson.M{"mid": mid}
	fmt.Println(where)
	err := mc.Table(t.Table()).Where(where).Find(&_list)
	if err != nil {
		return
	}
	return _list
}

type MovementHtml struct {
	// 最上方得提示信息
	ST string
	//作者信息
	Author struct {
		Name   string
		Avatar string
		Des    string
	}
	//多少人赞同
	Zan   int
	TextS struct {
		H     string
		Imags []string
	}
	Text string
}

func (t MovementCenter) ToConversion() (hs MovementHtml, err error) {
	var json []byte
	json, err = bson.MarshalJSON(t.V)
	if err != nil {
		return
	}
	fmt.Println(t.Type)
	switch t.Type {
	case MovementArticleSend:
		d := articleSend{}
		err = bson.Unmarshal(json, &d)
		if err != nil {
			return
		}
		fmt.Println(d)
		break
	case MovementQuestionSend:
		d := questionSend{}
		err = bson.UnmarshalJSON(json, &d)
		if err != nil {
			return
		}
		//直接组装html
		index := DataIndex{}
		mc.Table(index.Table()).Where(bson.M{"did": d.DID}).FindOne(&index)
		if index.DID == 0 {
			err = errors.New("提问丢失")
			return
		}
		hs.ST = "发布了问题"
		//获取数据详情
		mc.Table(index.InfoQuestion.Table()).Where(bson.M{"_id": index.ID}).FindOne(&index.InfoQuestion)
		if index.InfoQuestion.ID.Hex() == mc.Empty {
			err = errors.New("提问丢失")
			return
		}
		hs.TextS.H = index.InfoQuestion.Content
		hs.TextS.Imags = index.InfoQuestion.Imgs
		break
	}
	return
}
