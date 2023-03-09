package cache

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
)

var defaultCache *gcache.Cache

func InitCache(ctx context.Context) {
	var (
		initDefaultCache = false
		redisClient      *gredis.Redis
		err              error
	)
	redisClient = initRedisClient(ctx)
	if redisClient != nil {
		_, err = redisClient.Do(ctx, "SET", "Test", "test_value")
		if err == nil {
			defaultCache.SetAdapter(gcache.NewAdapterRedis(redisClient))
		} else {
			initDefaultCache = true
		}
	} else {
		initDefaultCache = true
	}

	if initDefaultCache {
		initDefaulCache()
	}
}

func initDefaulCache() {
	defaultCache = gcache.New()
}

func SetCache(ctx context.Context, key, value string) {
	defaultCache.Set(ctx, key, value, 60*60*3 * time.Second)
}

func GetCache(ctx context.Context, key string) (*gvar.Var, error) {
	return defaultCache.Get(ctx, key)
}