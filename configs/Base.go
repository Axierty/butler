package configs

import "github.com/astaxie/beego/config"

type Config struct {
	DbCommonConfig
	RedisCommonConfig
}

//数据库配置文件
type DbCommonConfig struct {
	DbHost     string
	DbPort     string
	DbDatabase string
	DbUsername string
	DbPassword string
	DbCharset  string
}

//redis配置文件
type RedisCommonConfig struct {
	RedisHost      string
	RedisPort      string
	RedisPass      string
	RedisMaxIdle   int
	RedisMaxActive int
}

var AppConfig *Config
var DbConfig *DbCommonConfig
var RedisConfig *RedisCommonConfig

func init() {
	AppConfig = &Config{}
	DbConfig = &DbCommonConfig{}
	RedisConfig = &RedisCommonConfig{}

	//初始化文件
	initConfig(".env")
}

func initConfig(file string) {
	conf, err := config.NewConfig("ini", file)
	if err != nil {
		panic(err)
	}
	initDbConfig(conf)
	initRedisConfig(conf)
}

func initDbConfig(conf config.Configer) {
	DbConfig.DbHost = conf.DefaultString("DB_HOST", "127.0.0.1")
	DbConfig.DbPort = conf.DefaultString("DB_PORT", "8080")
	DbConfig.DbDatabase = conf.DefaultString("DB_DATABASE", "test")
	DbConfig.DbUsername = conf.DefaultString("DB_USERNAME", "root")
	DbConfig.DbPassword = conf.DefaultString("DB_PASSWORD", "")
	DbConfig.DbCharset = conf.DefaultString("DB_CHARSET", "utf8mb4")
}

func initRedisConfig(conf config.Configer) {
	RedisConfig.RedisHost = conf.DefaultString("REDIS_HOST", "127.0.0.1")
	RedisConfig.RedisPort = conf.DefaultString("REDIS_PORT", "6357")
	RedisConfig.RedisPass = conf.DefaultString("REDIS_PASS", "")
	RedisConfig.RedisMaxIdle = conf.DefaultInt("REDIS_MAX_IDLE", 3)
	RedisConfig.RedisMaxActive = conf.DefaultInt("REDIS_MAX_ACTIVE", 5)
}
