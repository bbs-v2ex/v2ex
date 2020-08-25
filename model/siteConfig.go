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
	T                   string   `json:"t" bson:"t"`
	D                   string   `json:"d" bson:"d"`
	K                   string   `json:"k" bson:"k"`
	T_                  string   `json:"t_" bson:"t_"`
	NavigationHomeTitle string   `json:"navigation_home_title" bson:"navigation_home_title"`
	TitleDelimiter      string   `json:"title_delimiter" bson:"title_delimiter"`
	ICP                 string   `json:"icp" bson:"icp"`
	Activity            _____tdk `json:"activity" bson:"activity"`
	Question            _____tdk `json:"question" bson:"question"`
	Article             _____tdk `json:"article" bson:"article"`
	WX                  string   `json:"wx" bson:"wx"`
}
type _____tdk struct {
	T string `json:"t" bson:"t"`
	D string `json:"d" bson:"d"`
	K string `json:"k" bson:"k"`
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
