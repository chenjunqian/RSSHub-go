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
