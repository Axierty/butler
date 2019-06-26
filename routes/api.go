package routes

import (
	. "butler/controllers"
	. "butler/middlewares"
	"github.com/gin-gonic/gin"
)

func Api(Router *gin.Engine)  {

	Router.LoadHTMLGlob("resources/views/**/*")

	//默认页面
	Router.GET("/",Index.Index)

	//登录接口
	Router.POST("/login",Login.Login)

	//使用api 中间件
	api := Router.Group("/api", ApiToken())
	{
		api.GET("/update",Index.Update)
		api.GET("/test",Index.Test)
	}
}