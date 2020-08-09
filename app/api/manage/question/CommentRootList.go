package question

import (
	"fmt"
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/app/api"
	"v2ex/app/common"
	"v2ex/model"
)

/**
keep_self 保留自己
*/
func CommentRootList(did model.DIDTYPE, rid primitive.ObjectID, keep_self bool) []gin.H {
	//查询评论
	comment_list := []model.CommentQuestionRoot{}

	where := bson.M{"did": did}

	if rid.Hex() != mc.Empty {
		//首先查询一条数据出来
		_tmp_root := model.CommentQuestionRoot{}
		lt := "$lt"
		if keep_self {
			lt = "$lte"
		}
		where["_id"] = bson.M{lt: rid}
		mc.Table(_tmp_root.Table()).Where(bson.M{"_id": rid}).FindOne(&_tmp_root)
		if _tmp_root.ID.Hex() != mc.Empty {
			//fmt.Println(rid.Hex(),_tmp_root.ID.Hex())
			where["zan_len"] = bson.M{"$lte": _tmp_root.ZanLen}
		}
	}

	fmt.Println(where)
	mc.Table(model.CommentQuestionRoot{}.Table()).Where(where).Order(bson.M{"zan_len": -1, "_id": -1}).Limit(10).Find(&comment_list)
	//提取文本
	new_comment_list := []gin.H{}
	for k, _ := range comment_list {
		mc.Table(comment_list[k].Text.Table()).Where(bson.M{"_id": comment_list[k].ID}).FindOne(&comment_list[k].Text)
		if comment_list[k].Text.Text == "" {
			continue
		}
		//获取会员数据
		user_info := model.Member{}.GetUserInfo(comment_list[k].MID, true)
		new_comment_list = append(new_comment_list, gin.H{
			"user_info": gin.H{
				"name":   user_info.UserName,
				"avatar": common.Avatar(user_info.Avatar),
				"time":   c_code.StrTime(comment_list[k].Text.ReleaseTime),
				"mid":    comment_list[k].MID,
				"des":    user_info.More.Des,
			},
			"txt":      api.RestorePicture(comment_list[k].Text.Text, "", comment_list[k].Text.Img),
			"zan":      comment_list[k].ZanLen,
			"zan_user": comment_list[k].Text.Zan,
			"rc":       comment_list[k].RC,
			"_id":      comment_list[k].Text.ID,
		})
	}
	return new_comment_list
}
