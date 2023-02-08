package cache

import (
	"context"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gvar"
)

var defaultCache *gmap.Map

func InitCache() {
	initRedisClient()
	if redisClient == nil {
		initDefaulCache()
	}
}

func initDefaulCache() {
	defaultCache = gmap.New()
}

func SetCache(ctx context.Context, key, value string) {

	if redisClient != nil {
		redisClient.Do(ctx,"SET", key, value)
		redisClient.Do(ctx,"EXPIRE", key, 60*60*3)
	} else {
		if defaultCache == nil {
			initDefaulCache()
		}
		defaultCache.Set(key, value)
	}
}

func GetCache(ctx context.Context, key string) (*gvar.Var, error) {
	if redisClient != nil {
		return redisClient.Do(ctx, "GET", key)
	} else {
		var cacheV = gvar.New(defaultCache.Get(key))
		return cacheV, nil
	}
}
