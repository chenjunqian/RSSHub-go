package main

import (
	_ "rsshub/boot"
	_ "rsshub/router"

	"github.com/gogf/gf/v2/frame/g"
)

func main() {
  s := g.Server()
  s.SetIndexFolder(true)
  s.SetServerRoot("./template")
  s.AddStaticPath("/static","./template/static")
  g.View().SetPath("./template")
  s.Run()
}
