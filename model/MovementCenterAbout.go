package model

import (
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
