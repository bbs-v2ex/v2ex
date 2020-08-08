package view

import "github.com/gin-gonic/gin"

func Render(c *gin.Context, name string, _ht gin.H) {
	c.HTML(200, name, _ht)
}

type ViewError struct {
	Message string
}

func R404(c *gin.Context, view ViewError) {
	c.String(404, view.Message)
}
