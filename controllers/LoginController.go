package controllers

import (
	"butler/configs"
	"butler/libs"
	"butler/models"
	"fmt"
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
	err,user := userModel.CheckPassword(mobile,password)

	if err != nil {
		i.fail(c,configs.FAIL,"账号密码错误")
		return
	}

	token := userModel.GenerateToken()

	fmt.Println(user.Id)

	//获取redis连接句柄
	rc := libs.GetRedisInstance().Get()
	defer rc.Close()

	key := "redis#"+token
	//token一小时过期
	_, err2 := rc.Do("Set", key, user.Id,"EX", 60*60)
	if err2 != nil {
		fmt.Println(err2)
		i.fail(c,configs.FAIL,"redis设置失败")
		return
	}

	i.success(c,gin.H{"token": token})
}

