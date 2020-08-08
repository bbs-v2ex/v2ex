package model

import (
	"github.com/123456/c_code/mc"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SiteConfig struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id"`
	Key   int                `json:"key" bson:"key"`
	ValID interface{}        `json:"val" bson:"val"`
}

func (t SiteConfig) Table() string {
	return "site_config"
}

type SiteConfigSeo struct {
	T              string
	D              string
	K              string
	T_             string
	TitleDelimiter string
}

const _site_d_seo = 1

func (t SiteConfig) GetSeo() (sc SiteConfigSeo) {
	mc.Table(t.Table()).Where(bson.M{"key": _site_d_seo}).FindOne(&sc)
	return
}

func (t SiteConfig) SetSeo(sc SiteConfigSeo) error {
	err := mc.Table(t.Table()).Where(bson.M{"key": _site_d_seo}).UpdateOneIsEmptyNewInsert(&sc)
	return err
}
