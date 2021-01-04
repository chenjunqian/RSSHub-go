package service

import (
	"strings"

	"github.com/gogf/gf/frame/g"
)

func GetSiteCookies(siteName string) map[string]string {
	cookiesMap := make(map[string]string)
	cookiesString, err := g.Redis().DoVar("GET", siteName)
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
