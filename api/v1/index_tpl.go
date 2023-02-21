package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HomePageReq struct {
	g.Meta `path:"/" tags:"HomeTpl" method:"get" summary:"Home page tmplate request"`
}
type HomePageRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
