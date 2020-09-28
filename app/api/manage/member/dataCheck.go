package member

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/app/api"
	"v2ex/model"
)

func dataCheck(c *gin.Context) {
	mid := api.GetMID(c)
	//查询是否是管理员
	member := model.Member{}.GetUserInfo(mid, true)
	where := bson.M{"mid": mid}
	if member.MemberType == model.MemberTypeRoot {
		where = bson.M{}
	}
	//查询数据
	list := []model.DataCheck{}
	mc.Table(model.DataCheck{}.Table()).Where(where).Limit(100).Find(&list)
	l := []gin.H{}
	for _, v := range list {
		_one := gin.H{
			"time": c_code.StrTime(v.Itime),
		}
		_cz := []gin.H{
			{
				"t":    "查看",
				"u":    "data_check_view?id=" + v.ID.Hex(),
				"ajax": false,
			},
		}
		//if member.MemberType == model.MemberTypeRoot {
		//	_cz = append(_cz, []gin.H{
		//		{
		//			"t":    "通过",
		//			"u":    "data_check_view?id=" + v.ID.Hex() + "&p=pass",
		//			"ajax": false,
		//		},
		//		{
		//			"t":    "拒绝",
		//			"u":    "data_check_view?id=" + v.ID.Hex() + "&p=deny",
		//			"ajax": false,
		//		},
		//	}...)
		//}
		switch v.Type {
		case model.DataCheckTypeAddArticle:
			_one["tip"] = "发布文章"
			_one["title"] = v.D["title"]
			break
		}
		_one["cz"] = _cz
		l = append(l, _one)
	}
	c.JSON(200, c_code.V1GinSuccess(l))
}
