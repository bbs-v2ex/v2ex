package view

import (
	"bytes"
	"fmt"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func RenderGetContent(name string, _ht gin.H) string {
	buff := new(bytes.Buffer)
	err := ViewEngine.RenderWriter(buff, name, _ht)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	htmlContent := string(buff.Bytes())
	return htmlContent
}

func Html(c *gin.Context, htmlContent string) {
	c.Header("content-type", " text/html; charset=UTF-8")
	c.String(200, htmlContent)
}

func Render(c *gin.Context, name string, _ht gin.H) {
	_ht["navigation"] = setNavigation(c, _ht)
	c.HTML(200, name, _ht)
}

type ViewError struct {
	Message string
}

var ViewEngine = &ginview.ViewEngine{}

func R404(c *gin.Context, view ViewError) {
	c.String(404, view.Message)
}
