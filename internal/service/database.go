package service

import (
	"context"
	"fmt"
	"os"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var (
	dbInstance *gorm.DB
)

func InitDatabase(ctx context.Context) {
	databaseType, _ := g.Cfg().Get(ctx, "database.type")
	g.Log().Line().Infof(ctx, "database type is : %s \n", databaseType.String())
	setMySQLConfig(ctx)
}

func setMySQLConfig(ctx context.Context) {
	var (
		err         error
		dbConfig    gorm.Config
		dsn         string
		user        *gvar.Var
		password    *gvar.Var
		url         *gvar.Var
		dbName      *gvar.Var
		env         string
	)

	user, _ = g.Cfg().Get(ctx, "database.mysql.user")
	password, _ = g.Cfg().Get(ctx, "database.mysql.password")
	url, _ = g.Cfg().Get(ctx, "database.mysql.url")
	dbName, _ = g.Cfg().Get(ctx, "database.mysql.dbName")

	env = os.Getenv("env")
	if env == "dev" {
		dbConfig = gorm.Config{
			Logger: gormLogger.Default.LogMode(gormLogger.Info),
		}
		g.Log().Line().Info(ctx, "gorm is in dev mode")
	}
	dsn = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user.String(), password.String(), url.String(), dbName.String())
	dbInstance, err = gorm.Open(mysql.Open(dsn), &dbConfig)

	if err != nil {
		g.Log().Line().Error(ctx, err)
		panic(err)
	}
}

func GetDatabase() *gorm.DB {
	return dbInstance
}
