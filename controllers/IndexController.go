package controllers

import (
	"butler/models"
	"github.com/gin-gonic/gin"
)

var Index = &IndexController{}

type IndexController struct {
	BaseController
}

func (i *IndexController) Test(c *gin.Context) {
	i.success(c,gin.H{"message":"index"})
}

func (i *IndexController) Index(c *gin.Context) {

	userModel := &models.User{}

	user := userModel.Find(1)
	//fmt.Println(user)

	c.HTML(200, "index.html", gin.H{
		"title": "Test Title",
		"user" : user,
	})
}
