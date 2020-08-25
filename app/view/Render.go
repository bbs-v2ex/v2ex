package view

import (
	"bytes"
	"fmt"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"v2ex/app/nc"
	"v2ex/view_func"
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
	view_type, _ := c.Get("view_type")
	_ht["view_type"] = view_type
	if _ht["k"].(string) == "" {
		seoconfig := nc.GetSeoConfig()
		_ht["k"] = seoconfig.K
	}
	_ht["relesae_js"] = fmt.Sprintf("<script type=\"application/javascript\" src=\"%s\"></script>", view_func.ST("/js/release.js"))
	switch view_type {
	case "manage", "123":
		_ht["relesae_js"] = ""
		break
	}
	c.HTML(200, name, _ht)
}

type ViewError struct {
	Message string
}

var ViewEngine = &ginview.ViewEngine{}

func R404(c *gin.Context, view ViewError) {
	c.String(404, view.Message)
}
