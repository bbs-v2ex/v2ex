package config

//定义 系统基础设置
type SConfig struct {

	//当前系统得执行路径
	ExecPath string `toml:"-"`

	Run run `toml:"运行配置"`
	DB  db  `toml:"数据库配置"`
}

// 运行时得配置
type run struct {
	//调试模式 // true 不会编辑模板文件,会在命令行输出数据
	Debug bool `toml:"是否调试"`
	//如果设置 127.0.0.1 则需要设置方向代理才可方法
	LocaIP string `toml:"监听地址"`
	Port   int    `toml:"程序启动端口"`

	//临时上传目录
	TempUploadDir string `toml:"临时上传目录"`

	UploadServer string `toml:"图片服务器接口"`

	SiteMapUrlPreFix string `toml:"网站地图URL前缀"`
}

//数据量配置
type db struct {
	IP         string `toml:"IP地址"`
	Port       int    `toml:"端口"`
	DbName     string `toml:"数据库"`
	AuthSource string `toml:"安全验证"`
	UserName   string `toml:"账号"`
	PassWord   string `toml:"密码"`
}
