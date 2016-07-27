package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/zhuharev/streams"
	"github.com/zhuharev/streams/providers/twitch"
)

func main() {

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
	}

	app.Commands = []cli.Command{
		{

			Name:    "update",
			Aliases: []string{"up"},
			Usage:   "update all channels",
			Action:  update,
		},
	}

	app.Run(os.Args)

}
