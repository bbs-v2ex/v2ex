package nc

import "v2ex/model"

var seoconfig = model.SiteConfigSeo{}

func Init() {
	ReloadSeoConfig()
}

func GetSeoConfig() model.SiteConfigSeo {
	return seoconfig
}

func ReloadSeoConfig() {
	seoconfig = model.SiteConfig{}.GetSeo()
}
