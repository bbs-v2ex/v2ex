package member

import (
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"v2ex/model"
)

func data_check_view(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Query("id"))
	data_check := model.DataCheck{}
	mc.Table(data_check.Table()).Where(bson.M{"_id": id}).FindOne(&data_check)
	c.JSON(200, c_code.V1GinSuccess(data_check.D))
	return
}
