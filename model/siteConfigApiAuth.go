package model

import "github.com/123456/c_code/mc"
import "github.com/globalsign/mgo/bson"

const _site_api_auth = 2

type SiteConfigApiAuth struct {
	Register bool `json:"register" bson:"register"`
}

func (t SiteConfig) GetApiAuth() (sc SiteConfigApiAuth) {
	mc.Table(t.Table()).Where(bson.M{"key": _site_api_auth}).FindOne(&sc)
	return
}

func (t SiteConfig) SetApiAuth(sc SiteConfigApiAuth) error {
	err := mc.Table(t.Table()).Where(bson.M{"key": _site_api_auth}).UpdateOneIsEmptyNewInsert(&sc)
	return err
}
