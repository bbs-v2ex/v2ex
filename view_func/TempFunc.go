package view_func

import (
	"html/template"
	"strings"
)

func TempFunc() template.FuncMap {
	f := template.FuncMap{}
	f["html"] = func(s string) template.HTML {
		h := template.HTML(s)
		return h
	}
	f["st"] = func(urls ...string) string {
		u := strings.Join(urls, "/")
		return "/static/" + strings.TrimLeft(u, "/")
	}
	f["u"] = func(urls ...string) string {
		u := strings.Join(urls, "/")
		return "/" + strings.TrimLeft(u, "/")
	}
	return f
}
