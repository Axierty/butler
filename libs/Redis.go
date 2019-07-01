package libs

import (
	"butler/configs"
	"github.com/gomodule/redigo/redis"
	"time"
)

type RedisLib struct {
	RedisPool  *redis.Pool
	IsInit int
}

var RedisCore = &RedisLib{}

func GetRedisInstance() *redis.Pool {

	if RedisCore.IsInit == 0 {
		server := configs.RedisConfig.RedisHost + ":" + configs.RedisConfig.RedisPort
		redisPool := &redis.Pool{
			MaxIdle:     configs.RedisConfig.RedisMaxIdle,
			MaxActive:   configs.RedisConfig.RedisMaxActive,
			IdleTimeout: 240 * time.Second,
			Wait:        true,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", server, redis.DialPassword(configs.RedisConfig.RedisPass))
				if err != nil {
					return nil, err
				}
				return c, nil
			},
		}

		RedisCore.RedisPool = redisPool
		RedisCore.IsInit = 1
	}
	return RedisCore.RedisPool
}
