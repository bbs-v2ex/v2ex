package article

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/app/api"
	"v2ex/model"
)

func zan_del(c *gin.Context) {
	_f := _zan{}
	c.BindJSON(&_f)
	validator := api.VerifyValidator(_f)
	if validator != "" {
		result_json := c_code.V1GinError(101, validator)
		c.JSON(200, result_json)
		return
	}

	//查询数据
	comment_root := model.CommentRoot{}
	mc.Table(comment_root.Table()).Where(bson.M{"_id": _f.ID}).FindOne(&comment_root)
	if comment_root.ID.Hex() == mc.Empty {
		result_json := c_code.V1GinError(102, "未查询到数据")
		c.JSON(200, result_json)
		return
	}
	//查询数据
	comment_text := comment_root.Text
	mc.Table(comment_text.Table()).Where(bson.M{"_id": _f.ID}).FindOne(&comment_text)
	if comment_text.ID.Hex() == mc.Empty {
		result_json := c_code.V1GinError(102, "未查询到数据")
		c.JSON(200, result_json)
		return
	}
	mid := api.GetMID(c)
	//zan := []int(comment_text.Zan)
	zan := []int{}
	for _, v := range comment_text.Zan {
		zan = append(zan, int(v))
	}

	if !c_code.InArrayInt(int(mid), zan) {
		result_json := c_code.V1GinError(102, "未赞过")
		c.JSON(200, result_json)
		return
	}

	_zan := c_code.RemoveArrayInt(int(mid), zan)
	//更新字段
	err := mc.Table(comment_text.Table()).Where(bson.M{"_id": _f.ID}).UpdateOne(bson.M{"zan": _zan})
	if err != nil {
		result_json := c_code.V1GinError(103, "更新text表失败")
		c.JSON(200, result_json)
		return
	}
	err = mc.Table(model.CommentRoot{}.Table()).Where(bson.M{"_id": _f.ID}).UpdateOne(bson.M{"zan_len": len(_zan)})
	if err != nil {
		result_json := c_code.V1GinError(104, "更新root表失败")
		c.JSON(200, result_json)
		return
	}
	result_json := c_code.V1GinSuccess(_zan)
	//添加进通知中心
	model.Notice(mid, comment_root.MID).DelArticleZan(comment_root)
	c.JSON(200, result_json)
	return
}
