package member

import (
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/app/nc"
	"v2ex/app/view"
	"v2ex/app/view/controller"
	"v2ex/model"
)

func Index(c *gin.Context) {
	_ht := defaultData(c)

	//获取最近活跃的会员
	mids := []model.MIDTYPE{}
	_last_hd_member := model.MovementCenter{}.GetHuoYueMID(30)
	for _, v := range _last_hd_member {
		mids = append(mids, v["mid"].(model.MIDTYPE))
	}

	_ht["last_hd_member"] = _last_hd_member
	//获取新加入的会员

	_insert_hd_member := []model.Member{}
	insert_hd_member := []gin.H{}
	mc.Table(model.Member{}.Table()).Where(bson.M{"mid": bson.M{"$nin": mids}}).Order(bson.M{"_id": -1}).Limit(30).Find(&_insert_hd_member)
	for _, v := range _insert_hd_member {
		insert_hd_member = append(insert_hd_member, gin.H{
			"t":      v.UserName,
			"u":      model.UrlMember(v),
			"avatar": model.Avatar(v.Avatar),
		})
	}
	seoconfig := nc.GetSeoConfig()
	_ht["insert_hd_member"] = insert_hd_member
	_ht["t"] = controller.TitleJoin([]string{"会员列表"})
	_ht["d"] = "会员列表-" + seoconfig.D
	view.Render(c, "member/index", _ht)
}
