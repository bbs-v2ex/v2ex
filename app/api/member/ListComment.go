package member

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/model"
)

func ListComment(mid model.MIDTYPE, PID primitive.ObjectID) []model.MovementHtml {
	list := model.MovementCenter{}.Comment(mid)
	list_view := []model.MovementHtml{}
	for _, v := range list {
		hs, err := v.ToConversion()
		if err != nil {
			continue
		}
		//处理图片
		//hs.Text = api.RestorePicture(hs.TextS.H, "cover", hs.TextS.Imags)
		hs.Text = model.DesSplit(hs.TextS.H, 120)
		list_view = append(list_view, hs)
	}
	return list_view
}
