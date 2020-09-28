package root_api

import (
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/app/nc"
	"v2ex/model"
)

type _dataCheck struct {
	ID   primitive.ObjectID `json:"_id"`
	Type string             `json:"type"`
}

func dataCheck(c *gin.Context) {
	_f := _dataCheck{}
	c.BindJSON(&_f)
	_data := model.DataCheck{}
	mc.Table(_data.Table()).Where(bson.M{"_id": _f.ID}).FindOne(&_data)
	switch _f.Type {
	//修改状态为通过
	case "pass":
		//添加进文章表中
		c.JSON(200, nc.AddArticle(_data.MID, _data.D["title"].(string), _data.D["content"].(string), _data.Itime, true))
		break
	}
}
