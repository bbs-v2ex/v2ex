package model

import (
	"github.com/123456/c_code/mc"
	"time"
)
import "github.com/globalsign/mgo/bson"

const _site_update_sitemap = 3

type SiteConfigUpdateSiteMap struct {
	LastTime time.Time `json:"last_time" bson:"last_time" `
}

func (t SiteConfig) GetUpdateSiteMap() bool {
	sc := SiteConfigUpdateSiteMap{}
	mc.Table(t.Table()).Where(bson.M{"key": _site_update_sitemap}).FindOne(&sc)
	if sc.LastTime.Add(-1*time.Hour).Unix() > time.Now().Unix() {
		return false
	}
	return true
}

func (t SiteConfig) SetUpdateSiteMap(sc SiteConfigApiAuth) error {
	err := mc.Table(t.Table()).Where(bson.M{"key": _site_api_auth}).UpdateOneIsEmptyNewInsert(&sc)
	return err
}
