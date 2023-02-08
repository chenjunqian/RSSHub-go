package cache

import (
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

var redisClient *gredis.Redis

func initRedisClient() {
	redisClient = g.Redis()
}

func GetRedis() *gredis.Redis {
	return redisClient
}
