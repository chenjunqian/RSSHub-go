package service

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
)

func GetSiteCookies(ctx context.Context, siteName string) map[string]string {
	cookiesMap := make(map[string]string)
	cookiesString, err := g.Redis().Do(ctx, "GET", siteName)
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
