package controllers

import (
	"butler/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Index = &IndexController{}

type IndexController struct {
	BaseController
}

func (i *IndexController) Test(c *gin.Context) {

	//从token中解析出user_id 强转为int
	var _userId interface{}
	_userId,_ = c.Get("user_id")
	userId := _userId.(int)

	i.success(c,gin.H{"message":"index","user_id":userId})
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