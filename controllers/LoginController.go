package controllers

import (
	"butler/configs"
	"butler/models"
	"github.com/gin-gonic/gin"
)

var Login = &LoginController{}

type LoginController struct {
	BaseController
}

//登录接口
func (i *LoginController) Login(c *gin.Context) {

	mobile := c.PostForm("mobile")
	password := c.PostForm("password")

	if mobile == ""  || password == "" {
		i.fail(c,configs.FAIL,"账号密码错误")
		return
	}

	userModel := &models.User{}
	err,_ := userModel.CheckPassword(mobile,password)

	if err != nil {
		i.fail(c,configs.FAIL,"账号密码错误")
		return
	}

	token := userModel.GenerateToken()
	i.success(c,gin.H{"token": token})
}