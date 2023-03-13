package routers

import (
	"rsshub/internal/controller/rssapi/latexstudio"

	"github.com/gogf/gf/v2/net/ghttp"
)

func LatextStudioRouter(group *ghttp.RouterGroup) {
	group.GET("/article", latexstudio.Controller.GetArticle)
}
