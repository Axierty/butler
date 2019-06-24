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


func (i *IndexController) Update(c *gin.Context) {

	userModel := &models.User{}

	data := map[string]string{
		"name": "axxa",
	}

	dest := make(map[string]interface{})

	for k,v := range data {
		dest[k] = interface{}(v)
	}

	err := userModel.UpdateById(1,dest)
	if err != nil {
		i.fail(c,500,err.Error())
	}

	i.success(c,gin.H{"message":"edit sucess"})
}