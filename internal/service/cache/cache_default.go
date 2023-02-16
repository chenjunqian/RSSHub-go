package cache

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
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
	defaultCache.Set(ctx, key, value, 60*60*3)
}

func GetCache(ctx context.Context, key string) (*gvar.Var, error) {
	return defaultCache.Get(ctx, key)
}

func SetRouterInfoCache(ctx context.Context, key string, routerInfoList []RouterCategoryInfo) {
	defaultCache.Set(ctx, key, routerInfoList, 0)
}

func GetRouterInfoCache(ctx context.Context, key string) ([]RouterCategoryInfo, error) {
	var catagoryList []RouterCategoryInfo
	cacheV, err := defaultCache.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	err = gconv.Structs(cacheV, &catagoryList)
	return catagoryList, err
}
