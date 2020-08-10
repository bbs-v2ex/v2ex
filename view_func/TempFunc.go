package view_func

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"strings"
)

func TempFunc() template.FuncMap {
	f := template.FuncMap{}
	f["html"] = func(s string) template.HTML {
		h := template.HTML(s)
		return h
	}
	f["format"] = func(i interface{}) string {
		return fmt.Sprintf("%+v", i)
	}
	f["format2"] = func(conf interface{}) string {
		b, err := json.Marshal(conf)
		if err != nil {
			return fmt.Sprintf("%+v", conf)
		}
		var out bytes.Buffer
		err = json.Indent(&out, b, "", "    ")
		if err != nil {
			return fmt.Sprintf("%+v", conf)
		}
		return out.String()
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
