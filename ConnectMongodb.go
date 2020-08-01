package v2ex

import (
	"fmt"
	"github.com/123456/c_code"
	"github.com/123456/c_code/mc"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"v2ex/config"
)

func ConnectMongodb() {
	cg, err := config.LoadingConfigSourceFile()
	if err != nil {
		log.Fatal("加载配置文件失败", err)
		return
	}
	//生成 协议链接
	con := fmt.Sprintf("mongodb://%s:%d", cg.DB.IP, cg.DB.Port)
	//开始处理参数
	connect_options := options.Client().ApplyURI(con)
	//检测是否设置了 用户名及密码
	if cg.DB.PassWord != "" && cg.DB.AuthSource != "" {
		connect_options = options.Client().ApplyURI(con).SetAuth(options.Credential{
			AuthSource: cg.DB.AuthSource,
			Username:   cg.DB.UserName,
			Password:   cg.DB.PassWord,
		})
	}
	mgo_client, err := c_code.MongodbConnect(connect_options)
	if err != nil {
		log.Fatal(err.Error(), "Mongodb 链接失败")
	}
	mc.SetDB(mgo_client)
	//设置数据库
	mc.SetDBName(cg.DB.DbName)
}
