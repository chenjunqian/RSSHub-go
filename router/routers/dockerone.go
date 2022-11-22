package routers

import (
	"rsshub/app/api/rssapi/dockerone"

	"github.com/gogf/gf/v2/net/ghttp"
)

func DockerOneRouter(group *ghttp.RouterGroup) {
	group.GET("/recommand", dockerone.Controller.GetRecommand)
}