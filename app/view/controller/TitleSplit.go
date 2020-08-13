package controller

import (
	"github.com/123456/c_code"
	"strings"
	"v2ex/app/nc"
	"v2ex/until"
)

func TitleJoin(t []string) string {
	seoconfig := nc.GetSeoConfig()
	t = append(t, seoconfig.T_)
	return strings.Join(t, seoconfig.TitleDelimiter)
}

func KeywordJoin(d string) string {
	d = c_code.RemoveHtmlTag(d)
	return strings.Join(until.StopCIFilter(d), "，")
}

func DesJoin(d string) string {
	seoconfig := nc.GetSeoConfig()
	fgf := strings.TrimSpace(seoconfig.TitleDelimiter)
	d = d + fgf + seoconfig.D
	return d
}
