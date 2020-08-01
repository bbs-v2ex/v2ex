package config

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

var config_file_name = "000_config.toml"
var config SConfig

func CreateConfigFile() {
	ctoml := SConfig{
		DB: db{
			IP:       "127.0.0.1",
			UserName: "",
			PassWord: "",
		},
	}
	var buf bytes.Buffer
	toml.NewEncoder(&buf).Encode(ctoml)
	//写入配置文件
	ioutil.WriteFile(config_file_name, buf.Bytes(), os.ModePerm)
}

func LoadingConfigSourceFile() (config SConfig, err error) {
	_tmp := SConfig{}
	_, err = toml.DecodeFile(config_file_name, &_tmp)
	if err != nil {
		return
	}
	config = _tmp
	return
}
