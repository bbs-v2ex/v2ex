package member

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strconv"
	"v2ex/app/api"
	"v2ex/app/view"
	"v2ex/model"
)

func Index(c *gin.Context) {
	_ht := defaultData(c)
	mid, _ := strconv.Atoi(c.Param("mid"))

	u := fmt.Sprintf("/member/%d", mid)
	_member_mav := []gin.H{
		{
			"t":      "动态",
			"u":      u,
			"active": false,
		},
		{
			"t":      "提问",
			"u":      u + "/question",
			"active": false,
		},
		{
			"t":      "回答",
			"u":      u + "/comment",
			"active": false,
		},
		{
			"t":      "文章",
			"u":      u + "/article",
			"active": false,
		},
	}
	_type := c.Param("_type")
	_type_active := false
	for k, _ := range _member_mav {
		_f_last := filepath.Base(_member_mav[k]["u"].(string))
		if _f_last == _type {
			_type_active = true
			_member_mav[k]["active"] = true
			break
		}
	}

	if !_type_active {
		_member_mav[0]["active"] = true
	}
	_ht["_member_mav"] = _member_mav

	//查询是否存在此会员
	user_info := model.Member{}.GetUserInfo(model.MIDTYPE(mid), false)
	if user_info.MID == 0 {
		view.R404(c, view.ViewError{Message: "无此会员"})
		return
	}
	_ht["user_info"] = user_info
	list := model.MovementCenter{}.About(model.MIDTYPE(mid))
	list_view := []model.MovementHtml{}
	for _, v := range list {
		hs, err := v.ToConversion()
		if err != nil {
			continue
		}
		//处理图片
		hs.Text = api.RestorePicture(hs.TextS.H, "cover", hs.TextS.Imags)
		list_view = append(list_view, hs)
	}
	//_ht["dt"] = list_view

	view.Render(c, "member/user_home", _ht)
}
