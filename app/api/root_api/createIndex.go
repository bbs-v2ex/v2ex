package root_api

import (
	"context"
	"github.com/123456/c_code/mc"
	"github.com/gin-gonic/gin"
	"v2ex/model"
)

func createIndex(c *gin.Context) {
	mc.Coll(model.MovementCenter{}.Table()).Indexes().CreateMany(context.Background(), model.MovementCenter{}.IndexList())
	mc.Coll(model.CommentChild{}.Table()).Indexes().CreateMany(context.Background(), model.CommentChild{}.IndexList())
	mc.Coll(model.CommentRoot{}.Table()).Indexes().CreateMany(context.Background(), model.CommentRoot{}.IndexList())
	mc.Coll(model.CommentQuestionChild{}.Table()).Indexes().CreateMany(context.Background(), model.CommentQuestionChild{}.IndexList())
	mc.Coll(model.CommentQuestionRoot{}.Table()).Indexes().CreateMany(context.Background(), model.CommentQuestionRoot{}.IndexList())
	mc.Coll(model.DataIndex{}.Table()).Indexes().CreateMany(context.Background(), model.DataIndex{}.IndexList())
	mc.Coll(model.Member{}.Table()).Indexes().CreateMany(context.Background(), model.Member{}.IndexList())
	mc.Coll(model.MemberCollect{}.Table()).Indexes().CreateMany(context.Background(), model.MemberCollect{}.IndexList())
	mc.Coll(model.MemberToken{}.Table()).Indexes().CreateMany(context.Background(), model.MemberToken{}.IndexList())
}
