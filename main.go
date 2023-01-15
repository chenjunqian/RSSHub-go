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
  g.View().SetPath("./template")
  s.Run()
}
