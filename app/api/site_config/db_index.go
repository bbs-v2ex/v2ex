package site_config

import (
	"fmt"
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"reflect"
)

type _db_index_list struct {
	Key string `json:"key"`
}

func db_index(c *gin.Context) {

	l := []gin.H{}
	coll_names, err := mc.Client().ListCollectionNames(mc.TimeOut(), bson.M{})
	if err != nil {
		result_json := c_code.V1GinError(101, err.Error())
		c.JSON(200, result_json)
		return
	}
	for _, v := range coll_names {
		index_list, err := mc.Coll(v).Indexes().List(mc.TimeOut())
		if err != nil {
			continue
		}
		var r []bson.M
		err = index_list.All(mc.TimeOut(), &r)
		if err != nil {
			continue
		}
		istr := []interface{}{}
		for _, v2 := range r {
			fmt.Println(reflect.TypeOf(v2).String())
			istr = append(istr, v2["key"])
		}
		l = append(l, gin.H{
			"coll":  v,
			"index": istr,
		})
	}
	result_json := c_code.V1GinSuccess(l)
	c.JSON(200, result_json)
	return
}
