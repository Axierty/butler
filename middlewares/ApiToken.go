package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//api接口的中间件
func ApiToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		//c.Header("Access-Control-Allow-Origin", "test")
		//c.Set("name", "test")

		token := c.Request.Header.Get("token")

		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 422,
				"msg":  "api token 缺失",
				"data": "",
			})
			c.Abort()
		} else {
			c.Set("token", "11111")
		}

		//
		//for k,v :=range c.Request.Header {
		//	fmt.Println(k,v)
		//}

		c.Next()
	}
}
