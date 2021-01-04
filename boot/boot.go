package boot

import (
	_ "rsshub/packed"

	"github.com/gogf/gf/frame/g"
)

func init() {
	setCookiesToRedis()
}

func setCookiesToRedis() {
	cookiesMap := g.Cfg().GetMap("cookies")

	for key := range cookiesMap {
		g.Redis().DoVar("SET", key, cookiesMap[key])
	}
}
