package controller

import "github.com/gin-gonic/gin"

func DefaultData(c *gin.Context) (_ht gin.H) {
	_ht = gin.H{}
	_ht["_______API"] = "/api/manage"
	return
}

func defaultData(c *gin.Context) gin.H {
	return DefaultData(c)
}
