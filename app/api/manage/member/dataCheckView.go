package member

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/model"
)

func dataCheckView(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Query("id"))
	data_check := model.DataCheck{}
	where := bson.M{}
	if id.Hex() != mc.Empty {
		where = bson.M{"_id": id}
	}
	mc.Table(data_check.Table()).Where(where).FindOne(&data_check)

	result_data := data_check.D
	result_data["_id"] = data_check.ID

	c.JSON(200, c_code.V1GinSuccess(result_data))
	return
}
