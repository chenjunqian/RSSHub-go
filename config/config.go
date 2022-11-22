package config

import (
	"embed"
	"os"

	"github.com/gogf/gf/v2/encoding/gjson"
)

//go:embed *
var ConfigFS embed.FS

func GetConfig() *gjson.Json {
	var (
		env       string
		configStr []byte
		err       error
	)
	env = os.Getenv("env")
	if env == "dev" {
		configStr, err = ConfigFS.ReadFile("config.dev.json")
	} else {
		configStr, err = ConfigFS.ReadFile("config.json")
  }

	if err != nil {
		panic(err)
	}

	configJson := gjson.New(configStr)
	return configJson
}
