package nc

import "v2ex/model"

var seoconfig = model.SiteConfigSeo{}
var apiauth = model.SiteConfigApiAuth{}

func Init() {
	ReloadConfig()
}

func GetSeoConfig() model.SiteConfigSeo {
	return seoconfig
}
func GetApiAuth() model.SiteConfigApiAuth {
	return apiauth
}

func ReloadConfig() {
	seoconfig = model.SiteConfig{}.GetSeo()
	apiauth = model.SiteConfig{}.GetApiAuth()
}
