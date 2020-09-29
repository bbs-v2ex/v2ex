package article

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"v2ex/app/api"
	"v2ex/model"
)

type _comment_child_add struct {
	RID primitive.ObjectID `json:"rid" validate:"len=12" comment:"评论ID"`
	PID primitive.ObjectID `json:"pid"`
	Txt string             `json:"txt" validate:"min=10,max=1000" comment:"数据"`
}

func commentChildAdd(c *gin.Context) {
	//获取用户信息
	user_info := api.GetNowUserInfo(c)
	if user_info.MID == 0 {
		result_json := c_code.V1GinError(101, "非法操作")
		c.JSON(200, result_json)
		return
	}
	_f := _comment_child_add{}
	c.BindJSON(&_f)
	validator := api.VerifyValidator(_f)
	if validator != "" {
		result_json := c_code.V1GinError(102, validator)
		c.JSON(200, result_json)
		return
	}
	//删除html标签
	_f.Txt = c_code.RemoveHtmlTag(_f.Txt)

	//检测评论数据是否存在
	comment_root := model.CommentRoot{}
	mc.Table(comment_root.Table()).Where(bson.M{"_id": _f.RID}).FindOne(&comment_root)
	if comment_root.ID.Hex() == mc.Empty {
		result_json := c_code.V1GinError(103, "请勿乱传参")
		c.JSON(200, result_json)
		return
	}

	//通过 插入数据库
	comment_child := model.CommentChild{
		ID:     primitive.NewObjectID(),
		MID:    user_info.MID,
		RID:    _f.RID,
		PID:    _f.PID,
		ZanLen: 0,
	}
	err := mc.Table(comment_child.Table()).Insert(comment_child)
	if err != nil {
		result_json := c_code.V1GinError(104, "写入失败")
		c.JSON(200, result_json)
		return
	}

	//写入数据存储表
	comment_text := model.CommentText{
		ID:          comment_child.ID,
		Text:        _f.Txt,
		Zan:         nil,
		Img:         nil,
		ReleaseTime: time.Now(),
	}
	//写进数据表中
	err = mc.Table(comment_text.Table()).Insert(comment_text)

	if err != nil {
		result_json := c_code.V1GinError(105, "写入失败")
		mc.Table(comment_root.Table()).Where(bson.M{"_id": comment_child.ID}).DelOne()
		c.JSON(200, result_json)
		return
	}
	//主频率 RC 字段加一
	mc.Table(comment_root.Table()).Where(bson.M{"_id": _f.RID}).FieldAddOrDel("rc", +1)
	result_json := c_code.V1GinSuccess("", "评论成功")
	c.JSON(200, result_json)
}
