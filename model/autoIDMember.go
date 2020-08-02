package model

import (
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (t AutoID) MemberID() (id MIDTYPE, e error) {
	_a := &AutoID{}
	mc.Table(t.Table()).Where(bson.M{"t_id": AutoIDMember}).FindOne(&_a)
	//如果数据没有

	//if _a.TableID == 0 {
	//	e = errors.New("数据查询失败")
	//	return
	//}
	if _a.ValID <= 1 {
		mc.Table(t.Table()).Insert(&AutoID{
			ID:      primitive.NewObjectID(),
			TableID: AutoIDMember,
			ValID:   1,
		})
		id = 1
		return
	}
	id = MIDTYPE(_a.ValID)
	return
}
func (t AutoID) MemberAdd() (e error) {
	return mc.Table(t.Table()).Where(bson.M{"t_id": AutoIDMember}).FieldAddOrDel("val", +1)
}
