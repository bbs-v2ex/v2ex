package member

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/model"
)

func ListDynamic(mid model.MIDTYPE, PID primitive.ObjectID) []model.MovementHtml {
	list := model.MovementCenter{}.About(mid)
	list_view := []model.MovementHtml{}
	for _, v := range list {
		hs, err := v.ToConversion()
		if err != nil {
			continue
		}
		//处理图片
		//hs.Text = api.RestorePicture(hs.TextS.H, "cover", hs.TextS.Imags)

		hs.Text = model.DesSplit(hs.TextS.H, 120)
		hs.TextS.H = ""
		list_view = append(list_view, hs)
	}
	return list_view
}
