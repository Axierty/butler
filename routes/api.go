package routes

import (
	. "butler/controllers"
	"github.com/gin-gonic/gin"
)

func Api(Router *gin.Engine)  {
	Router.GET("/test",Index.Test)
}