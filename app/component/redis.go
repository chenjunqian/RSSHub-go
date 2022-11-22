package component

import (
	"rsshub/config"

	"github.com/gogf/gf/v2/database/gredis"
)

var redisClient *gredis.Redis

func InitRedisClient() {

	var (
		redisLink string
		db        int
		password  string
	)

	redisLink = config.GetConfig().Get("redis.default.url").String()
	db = config.GetConfig().Get("redis.default.db").Int()
	password = config.GetConfig().Get("redis.default.password").String()
	config := gredis.Config{
		Address: redisLink,
		Db:      db,
		Pass:    password,
	}

	gredis.SetConfig(&config)
	redisClient = gredis.Instance()

}

func GetRedis() *gredis.Redis {
	return redisClient
}
