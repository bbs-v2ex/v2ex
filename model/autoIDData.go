package model

import (
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
)

func (t AutoID) DataID() (id int64, e error) {
	_a := &AutoID{}
	mc.Table(t.Table()).Where(bson.M{"t_id": AutoIDData}).FindOne(_a)
	//如果数据没有

	//if _a.TableID == 0 {
	//	e = errors.New("数据查询失败")
	//	return
	//}
	if _a.ValID <= 1 {
		mc.Table(t.Table()).Insert(&AutoID{
			TableID: AutoIDData,
			ValID:   1,
		})
		id = 1
		return
	}
	id = _a.ValID
	return
}
func (t AutoID) DataAdd() (e error) {
	return mc.Table(t.Table()).Where(bson.M{"t_id": AutoIDData}).FieldAddOrDel("val", +1)
}
