package main

import (
	"github.com/urfave/cli"
)

func Init(c *cli.Context) error {
	e := NewConfigContext(c)
	if e != nil {
		return e
	}
	e = InitDb()
	if e != nil {
		return e
	}
	return nil
}
