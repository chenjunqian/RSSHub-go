package boot

import (
	"context"
	"github.com/gogf/gf/os/genv"
	"github.com/olivere/elastic/v7"
	"rsshub/app/cronJob"
	_ "rsshub/packed"

	"github.com/gogf/gf/frame/g"
)

var esClient *elastic.Client
var esContext context.Context

func init() {
	configEvn := genv.Get("GF_GCFG_FILE", "")
	if configEvn != "" {
		g.Cfg().SetFileName(configEvn)
	}
	//app 相关配置
	s := g.Server()
	//GF相关配置 Web Server配置
	g.Log().Stack(true)
	s.SetErrorLogEnabled(true)
	s.SetAccessLogEnabled(true)
	setCookiesToRedis()
	initES()
	cronJob.RegisterJob()
}

func setCookiesToRedis() {
	cookiesMap := g.Cfg().GetMap("cookies")

	for key := range cookiesMap {
		g.Redis().DoVar("SET", key, cookiesMap[key])
	}
}

func initES() {
	g.Log().Info("init elastic search")
	host := g.Cfg().GetString("elasticsearch.host")
	port := g.Cfg().GetString("elasticsearch.port")
	username := g.Cfg().GetString("elasticsearch.user")
	password := g.Cfg().GetString("elasticsearch.pass")
	url := host + ":" + port
	esContext = context.Background()
	var err error
	esClient, err = elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false), elastic.SetHealthcheck(false), elastic.SetBasicAuth(username, password))
	if err != nil {
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := esClient.Ping(url).Do(esContext)
	if err != nil {
		// Handle error
		panic(err)
	}
	g.Log().Infof("Elasticsearch returned with code %d and version %s", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esVersion, err := esClient.ElasticsearchVersion(url)
	if err != nil {
		// Handle error
		panic(err)
	}
	g.Log().Infof("Elasticsearch version %s", esVersion)
}

func GetESClient() *elastic.Client {
	return esClient
}

func GetESContext() context.Context {
	return esContext
}
