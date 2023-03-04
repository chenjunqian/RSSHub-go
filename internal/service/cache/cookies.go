package cache

import (
	"context"
	"errors"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
)

func GetSiteCookies(ctx context.Context, siteName string) map[string]string {
	var (
		err           error
		cookiesMap    g.MapStrStr
		cookiesString *g.Var
	)
	cookiesMap = make(map[string]string)
	cookiesString, err = defaultCache.Get(ctx, siteName)
	if err == nil {
		cookieStringArray := strings.Split(cookiesString.String(), ";")
		for _, item := range cookieStringArray {
			key := strings.Split(item, "=")[0]
			value := strings.Split(item, "=")[1]
			cookiesMap[key] = value
		}
	}

	return cookiesMap
}

func InitSiteCookies(ctx context.Context) error {
	var (
		err       error
		cookies   *g.Var
		cookieMap g.MapStrStr
	)
	cookies, err = g.Cfg().Get(ctx, "cookies")
	if err != nil {
		g.Log().Error(ctx, "Init site cookies failed : \n", err)
		return err
	}
	cookieMap = cookies.MapStrStr()
	if len(cookieMap) == 0 {
		return errors.New("Site cookie is empty")
	}

	for k := range cookieMap {
		SetCache(ctx, k, cookieMap[k])
	}

	return nil
}
