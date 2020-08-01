package config

//定义 系统基础设置
type SConfig struct {
	DB db `toml:"数据库配置"`
}

//数据量配置
type db struct {
	IP       string `toml:"IP地址"`
	UserName string `toml:"账号"`
	PassWord string `toml:"密码"`
}
