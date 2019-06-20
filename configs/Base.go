package configs

import "github.com/astaxie/beego/config"

type Config struct {
	DbCommonConfig
}

type DbCommonConfig struct {
	DbHost     string
	DbPort     string
	DbDatabase string
	DbUsername string
	DbPassword string
	DbCharset  string
}


var AppConfig *Config
var DbConfig *DbCommonConfig

func init() {
	AppConfig = &Config{}
	DbConfig = &DbCommonConfig{}

	//初始化文件
	InitConfig(".env")
}


func InitConfig(file string) {
	conf, err := config.NewConfig("ini", file)
	if err != nil {
		panic(err)
	}
	DbConfig.DbHost = conf.DefaultString("DB_HOST", "127.0.0.1")
	DbConfig.DbPort = conf.DefaultString("DB_PORT", "8080")
	DbConfig.DbDatabase = conf.DefaultString("DB_DATABASE", "test")
	DbConfig.DbUsername = conf.DefaultString("DB_USERNAME", "root")
	DbConfig.DbPassword = conf.DefaultString("DB_PASSWORD", "")
	DbConfig.DbCharset = conf.DefaultString("DB_CHARSET", "utf8mb4")

}