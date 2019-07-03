package middlewares

import (
	"butler/libs"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
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

			//获取redis连接句柄
			rc := libs.GetRedisInstance().Get()
			defer rc.Close()

			key := "redis#" + token
			userId, err := redis.Int(rc.Do("Get", key))

			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 500,
					"msg":  "redis获取失败",
					"data": "",
				})
				c.Abort()
			}

			//中间件中传递解析token 的用户id
			c.Set("user_id", userId)
		}

		//
		//for k,v :=range c.Request.Header {
		//	fmt.Println(k,v)
		//}

		c.Next()
	}
}
