package view_func

import "html/template"

func TempFunc() template.FuncMap {
	f := template.FuncMap{}
	f["html"] = func(s string) template.HTML {
		h := template.HTML(s)
		return h
	}
	return f
}
