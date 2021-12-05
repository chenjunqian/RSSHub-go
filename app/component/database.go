package component

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcfg"
)

var (
	databaseInstance gdb.DB
)

func InitDatabase() {
	var err error
	setDatabaseConfig()
	databaseInstance, err = gdb.New("default")
	if err != nil {
		g.Log().Error(err)
		panic(err)
	}
}

func setDatabaseConfig() {
	var (
		host   string
		port   string
		user   string
		pass   string
		name   string
		debug  bool
		dbType string
		config *gcfg.Config
	)
	config = g.Cfg()
	host = config.GetString("database.default.0.host")
	port = config.GetString("database.default.0.port")
	user = config.GetString("database.default.0.user")
	name = config.GetString("database.default.0.name")
	pass = config.GetString("database.default.0.pass")
	debug = config.GetBool("database.default.0.debug")
	dbType = config.GetString("database.default.0.type")
	gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			gdb.ConfigNode{
				Host:   host,
				Port:   port,
				User:   user,
				Pass:   pass,
				Name:   name,
				Debug:  debug,
				Type:   dbType,
				Weight: 100,
			},
		},
	})
}

func GetDatabase() gdb.DB {
	return databaseInstance
}
