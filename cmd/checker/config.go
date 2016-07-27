package main

import (
	"github.com/urfave/cli"
	"gopkg.in/ini.v1"
)

var (
	Config struct {
		Db      Db
		DevMode bool
	}

	iniFile *ini.File
)

type Db struct {
	Driver   string
	User     string
	Password string
	Host     string
	Dbname   string
	Path     string

	// for ql
	IdName string
}

func NewConfigContext(c *cli.Context) (e error) {
	iniFile, e = ini.Load(c.String("config"))
	if e != nil {
		return e
	}
	iniFile.NameMapper = ini.TitleUnderscore
	e = iniFile.MapTo(&Config)
	if e != nil {
		return e
	}

	if c.GlobalString("mode") == DEV {
		Config.DevMode = true
	}

	if Config.DevMode {
		e = iniFile.Section("devdb").MapTo(&Config.Db)
		if e != nil {
			return e
		}
	}

	return nil
}
