package main

import (
	"fmt"

	_ "github.com/go-xorm/ql"
	_ "github.com/lunny/ql/driver"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-xorm/xorm"

	"github.com/zhuharev/streams"
	//"github.com/zhuharev/streams/providers/twitch"
)

var (
	x      *xorm.Engine
	tables = []interface{}{
		new(streams.Channel),
		new(streams.Stream),
	}
)

func InitDb() (e error) {
	x, e = newEngine()
	if e != nil {
		return e
	}
	if Config.DevMode {
		x.ShowDebug = true
		x.ShowErr = true
		x.ShowWarn = true
		x.ShowInfo = true
		x.ShowSQL = true
	}

	// not handle this error, ql not support alter table
	e = x.Sync2(tables...)
	if e != nil {
		if !Config.DevMode {
			return e
		}
	}

	return nil
}

func newEngine() (*xorm.Engine, error) {
	var (
		conn string
	)

	Config.Db.IdName = "id"

	switch Config.Db.Driver {
	case "mysql":
		conn = fmt.Sprintf("%s:%s@/%s", Config.Db.User, Config.Db.Password, Config.Db.Dbname)
	case "ql":
		conn = Config.Db.Path
		Config.Db.IdName = "id()"
	}
	x, e := xorm.NewEngine(Config.Db.Driver, conn)
	return x, e
}
