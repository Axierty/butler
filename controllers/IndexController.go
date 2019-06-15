package controllers

import "github.com/gin-gonic/gin"

var Index = &IndexController{}

type IndexController struct {
	BaseController
}

func (i *IndexController) Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "index",
	})
}
