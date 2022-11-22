package boot

import (
	"context"
	"rsshub/app/component"
	"rsshub/config"
)

func init() {
	component.InitLogger()
	component.InitRedisClient()
	setCookiesToRedis()
}

func setCookiesToRedis() {
	ctx := context.Background()
	cookiesMap := config.GetConfig().Get("cookies").Map()
	for key := range cookiesMap {
		component.GetRedis().Do(ctx, "SET", key, cookiesMap[key])
	}
}
