package config

import (
	"embed"
	"os"

	"github.com/gogf/gf/v2/encoding/gjson"
)

//go:embed *
var ConfigFS embed.FS

var ConfigJson *gjson.Json

func GetConfig() *gjson.Json {
	var (
		env         string
		embedConfig string
		configStr   []byte
		err         error
	)
	if !ConfigJson.IsNil() {
		return ConfigJson
	}
	env = os.Getenv("env")
	embedConfig = os.Getenv("embedConfig")
	if embedConfig == "true" {
		if env == "dev" {
			configStr, err = ConfigFS.ReadFile("config.dev.json")
		} else {
			configStr, err = ConfigFS.ReadFile("config.json")
		}
	} else {
		if env == "dev" {
			ConfigJson, err = gjson.Load("./config/config.dev.json")
		} else {
			ConfigJson, err = gjson.Load("./config/config.json")
		}
		return ConfigJson
	}

	if err != nil {
		panic(err)
	}

	ConfigJson = gjson.New(configStr)
	return ConfigJson
}
