package models

import (
	"github.com/jinzhu/gorm"
)

// DB连接
var DB *gorm.DB

func BaseInit()  {
	initDB()
}

func initDB() {

}