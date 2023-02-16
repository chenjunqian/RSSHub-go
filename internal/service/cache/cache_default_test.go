package cache

import (
	"testing"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_initDefaulCache(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		initDefaulCache()
		if defaultCache == nil {
			t.Fatal("default cache init failed")
		}
	})
}

func TestSetGetCache(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			key    string
			value  string
			ctx    = gctx.New()
			err    error
			cacheV *gvar.Var
		)
		key = "test_key"
		value = "test_value"
		initDefaulCache()
		SetCache(ctx, key, value)
		cacheV, err = GetCache(ctx, key)
		if err != nil {
			t.Fatal("get cache failed : " + err.Error())
		}

		if cacheV == nil {
			t.Fatal("the value from cache is nil")
		}

		if cacheV.String() != value {
			t.Fatal("the value from cache is not match the input value")
		}
	})

}

func TestGetRouterInfoCache(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx       = gctx.New()
			key       string
			value     []RouterCategoryInfo
			valueItem RouterCategoryInfo
		)
		InitCache(ctx)
		key = "test_key"
		value = make([]RouterCategoryInfo, 0)
		valueItem = RouterCategoryInfo{
			Router:   "test_router",
			Category: "test_catagory",
		}
		value = append(value, valueItem)
		SetRouterInfoCache(ctx, key, value)
		cacheV, err := GetRouterInfoCache(ctx, key)
		if err != nil {	
			t.Fatal("get router catagory from cache failed: ", err)
		}

		if cacheV == nil {
			t.Fatal("the value get from cache is nil")
		}

		if len(cacheV) == 0 {
			t.Fatal("the router catagorey list from cache length is 0")
		}
	})

}
