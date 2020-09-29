package root_api

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"time"
	"v2ex/model"
)

type _search struct {
	MID   model.MIDTYPE `json:"mid"`
	Title string        `json:"title"`
	Type  string        `json:"type"`
}

func dataSearch(c *gin.Context) {
	_f := _search{}
	c.BindJSON(&_f)
	where := bson.M{}
	if _f.MID != 0 {
		where["mid"] = _f.MID
	}
	if _f.Type == "article" {
		where["d_type"] = model.DTYPEArticle
	}
	if _f.Type == "question" {
		where["d_type"] = model.DTYPEQuestion
	}
	if _f.Title != "" {
		where["t"] = bson.M{"$regex": _f.Title}
	}
	index_list := []model.DataIndex{}
	index_list_2 := []gin.H{}
	mc.Table(model.DataIndex{}.Table()).Where(where).Limit(25).Find(&index_list)
	for _, v := range index_list {

		d := gin.H{
			"_id":   v.ID,
			"title": v.T,
		}
		switch v.DTYPE {
		case model.DTYPEArticle:
			d["type"] = "文章"
			d["u"] = model.UrlArticle(v)
			break
		case model.DTYPEQuestion:
			d["type"] = "问题"
			d["u"] = model.UrlQuestion(v)
			break
		}
		d["time"] = c_code.StrTime(time.Unix(v.CT, 0))

		index_list_2 = append(index_list_2, d)
	}
	c.JSON(200, c_code.V1GinSuccess(index_list_2))
}
