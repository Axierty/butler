package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"butler/configs"
)

// DB连接
var DB *gorm.DB
var dbConfig = configs.DbConfig

func BaseInit()  {
	initDB()
}

func initDB() {

	url := dbConfig.DbUsername + ":" + dbConfig.DbPassword + "@tcp(" + dbConfig.DbHost + ":" + dbConfig.DbPort + ")/" + dbConfig.DbDatabase +
		"?charset=" + dbConfig.DbCharset

	var err error
	DB, err = gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	fmt.Println(DB)
}