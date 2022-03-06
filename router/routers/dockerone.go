package routers

import (
	"rsshub/app/api/rssapi/dockerone"

	"github.com/gogf/gf/net/ghttp"
)

func DockerOneRouter(group *ghttp.RouterGroup) {
	group.GET("/recommand", dockerone.Controller.GetRecommand)
}