package edit

import (
	"fmt"
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"v2ex/app/api"
	"v2ex/model"
)

func editpost(c *gin.Context) {
	_f := _edit_result{}
	err := c.Bind(&_f)
	if err != nil {
		result_json := c_code.V1GinError(100, err.Error())
		c.JSON(200, result_json)
		return
	}
	mid := api.GetMID(c)
	switch _f.Type {
	case "article":
		data_index := model.DataIndex{}
		mc.Table(data_index.Table()).Where(bson.M{"did": _f.DID, "mid": mid}).FindOne(&data_index)
		if data_index.MID != mid {
			result_json := c_code.V1GinError(101, "请勿乱传参")
			c.JSON(200, result_json)
			return
		}
		mc.Table(data_index.Table()).Where(bson.M{"_id": data_index.ID}).UpdateOne(bson.M{"t": _f.Title})
		content, imgs, err := api.SeparatePicture(_f.Content)
		if err != nil {
			result_json := c_code.V1GinError(102, "请勿乱传参")
			c.JSON(200, result_json)
			return
		}
		mc.Table(data_index.InfoArticle.Table()).Where(bson.M{"_id": data_index.ID}).UpdateOne(bson.M{"content": content, "imgs": imgs})

		result_json := c_code.V1GinSuccess("修改成功", "", fmt.Sprintf("/%s/%d", model.UrlTagArticle, _f.DID))
		c.JSON(200, result_json)
		return
	}

}
