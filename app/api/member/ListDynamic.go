package member

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/model"
)

func ListDynamic(mid model.MIDTYPE, PID primitive.ObjectID) []model.MovementHtml {
	list := model.MovementCenter{WID: PID}.About(mid)
	list_view := []model.MovementHtml{}
	for _, v := range list {
		hs, err := v.ToConversion()
		if err != nil {
			fmt.Println(v.ID.Hex())
			fmt.Println(err)
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
