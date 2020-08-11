package model

import (
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
	if mid == 0 {
		where = bson.M{}
	}
	if t.WID.Hex() != mc.Empty {
		where["_id"] = bson.M{"$lt": t.WID}
	}
	//fmt.Println(where)
	err := mc.Table(t.Table()).Where(where).Order(bson.M{"_id": -1}).Limit(10).Find(&_list)
	if err != nil {
		return
	}
	return _list
}

func (t MovementCenter) Comment(mid MIDTYPE) (list []MovementCenter) {
	_list := []MovementCenter{}

	//where := bson.M{
	//	"$or": []bson.M{
	//		{"mid": mid},
	//		//{"m2id": mid},
	//	},
	//}
	where := bson.M{"mid": mid, "type": MovementQuestionCommentRoot}
	//fmt.Println(where)
	err := mc.Table(t.Table()).Where(where).Order(bson.M{"_id": -1}).Limit(10).Find(&_list)
	if err != nil {
		return
	}
	return _list
}
