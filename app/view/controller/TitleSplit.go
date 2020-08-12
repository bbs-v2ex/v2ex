package controller

import (
	"strings"
	"v2ex/app/nc"
)

func TitleJoin(t []string) string {
	seoconfig := nc.GetSeoConfig()
	t = append(t, seoconfig.T_)
	return strings.Join(t, seoconfig.TitleDelimiter)
}

func DesJoin(d string) string {
	seoconfig := nc.GetSeoConfig()
	fgf := strings.TrimSpace(seoconfig.TitleDelimiter)
	d = d + fgf + seoconfig.D
	return d
}
