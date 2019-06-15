package main

import (
	"butler/models"
	"butler/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

//gin的app容器
var app = gin.Default()

func init() {
	fmt.Println("app init")

	//初始化orm
	models.BaseInit()

	//注册路由
	routes.Api(app)
}

//
func main() {

	app.Run(":8081")
}
