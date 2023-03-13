package cache

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
)

func initRedisConfig(ctx context.Context) *gredis.Config {
	var (
		err         error
		redisConfig = &gredis.Config{}
		redisAddr   *gvar.Var
		db          *gvar.Var
		pass        *gvar.Var
	)

	redisAddr, err = gcfg.Instance().Get(ctx, "redis.default.address")
	if err != nil {
		g.Log().Panic(ctx, err)
	}
	db, err = gcfg.Instance().Get(ctx, "redis.default.db")
	if err != nil {
		g.Log().Panic(ctx, err)
	}
	pass, err = gcfg.Instance().Get(ctx, "redis.default.pass")
	if err != nil {
		g.Log().Panic(ctx, err)
	}

	redisConfig.Address = redisAddr.String()
	redisConfig.Db = db.Int()
	redisConfig.Pass = pass.String()

	return redisConfig
}

func initRedisClient(ctx context.Context) *gredis.Redis {
	var (
		err         error
		redisClient *gredis.Redis
	)

	redisClient, err = gredis.New(initRedisConfig(ctx))
	if err != nil {
		g.Log().Panic(ctx, err)
	}

	return redisClient
}
