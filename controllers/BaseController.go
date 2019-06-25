package controllers

import (
	"butler/configs"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct{}

func (*BaseController) success(c *gin.Context, data map[string]interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": configs.SUCCESS,
		"msg":  "请求成功",
		"data": data,
	})
}


func (*BaseController) fail(c *gin.Context,  code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": "",
	})
}