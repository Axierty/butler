package routes

import (
	. "butler/controllers"
	"github.com/gin-gonic/gin"
)

func Api(Router *gin.Engine)  {

	Router.LoadHTMLGlob("resources/views/**/*")


	Router.GET("/test",Index.Test)

	Router.GET("/",Index.Index)

	Router.GET("/update",Index.Update)


	Router.POST("/login",Login.Login)
}