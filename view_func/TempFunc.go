package view_func

import (
	"encoding/json"
	"html/template"
	"strings"
)

func TempFunc() template.FuncMap {
	f := template.FuncMap{}
	f["html"] = func(s string) template.HTML {
		h := template.HTML(s)
		return h
	}
	f["json"] = func(s interface{}) string {
		marshal, _ := json.Marshal(s)
		return string(marshal)
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
