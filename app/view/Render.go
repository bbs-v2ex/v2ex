package view

import "github.com/gin-gonic/gin"

func Render(c *gin.Context, name string, _ht gin.H) {
	c.HTML(200, name, _ht)
}
