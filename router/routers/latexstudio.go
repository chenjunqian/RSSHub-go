package routers

import (
	"rsshub/app/api/rssapi/latexstudio"

	"github.com/gogf/gf/net/ghttp"
)

func LatextStudioRouter(group *ghttp.RouterGroup) {
	group.GET("/article", latexstudio.Controller.GetArticle)
}
