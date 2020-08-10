package member

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"v2ex/app/api"
	"v2ex/app/view"
	"v2ex/model"
)

func Index(c *gin.Context) {
	_ht := defaultData(c)
	mid, _ := strconv.Atoi(c.Param("mid"))
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
	_ht["dt"] = list_view
	view.Render(c, "member/user_home", _ht)
}
