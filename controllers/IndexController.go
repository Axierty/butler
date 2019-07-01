package controllers

import (
	"butler/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Index = &IndexController{}

type IndexController struct {
	BaseController
}

func (i *IndexController) Test(c *gin.Context) {

	user_id,_ := c.Get("user_id")
	fmt.Println("ssss",user_id)

	i.success(c,gin.H{"message":"index"})
}

func (i *IndexController) Index(c *gin.Context) {

	userModel := &models.User{}

	user := userModel.Find(1)

	users := userModel.FindAll()
	//fmt.Println(user)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Test Title",
		"user" : user,
		"users" : users,
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
		return
	}

	i.success(c,gin.H{"message":"edit sucess"})
}