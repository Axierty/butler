package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB连接
var DB *gorm.DB

func BaseInit()  {
	initDB()
}

func initDB() {
	//var err error
	//DB, err = gorm.Open("mysql", "root:root/short?charset=utf8&parseTime=True&loc=Local")
	//if err != nil {
	//	panic(err)
	//}
}