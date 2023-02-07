package routers

import (
	"rsshub/internal/controller/rssapi/dockerone"

	"github.com/gogf/gf/v2/net/ghttp"
)

func DockerOneRouter(group *ghttp.RouterGroup) {
	group.GET("/recommand", dockerone.Controller.GetRecommand)
}