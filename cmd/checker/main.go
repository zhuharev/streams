package main

import (
	//"fmt"
	"os"

	"github.com/urfave/cli"
	//"github.com/zhuharev/streams"
	//"github.com/zhuharev/streams/providers/twitch"
)

const (
	devmode = "devmode"

	DEV  = "dev"
	PROD = "prod"
)

func main() {

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "mode, m",
			Value: PROD,
			Usage: "Production or developerment mode",
		},
		cli.BoolFlag{
			Name:  devmode + ", d",
			Usage: "use flag for reading config from dev section",
		},
		stringFlag("config, c", "app.ini", "add config"),
	}

	app.Commands = []cli.Command{
		{
			Name:    "update",
			Aliases: []string{"up"},
			Usage:   "update all channels",
			Action:  update,
			Flags:   []cli.Flag{
			//intFlag("port, p", 3000, "Temporary port number to prevent conflict"),
			//stringFlag("config, c", "conf/app.ini", "Configuration file path (default ./conf/app.ini)"),
			//stringFlag("mode", "prod", "Running mode"),
			},
		},
		{
			Name:   "add",
			Usage:  "add new channel",
			Action: add,
			Flags: []cli.Flag{
				stringFlag("url", "", "Add channel with `URL`"),
				stringFlag("service, s", "", "Add channel with `SERVICE`, user required"),
				stringFlag("username, u", "", "Add channel with `USER`, service required"),
			},
		},
		{
			Name:   "status",
			Usage:  "current channel status",
			Action: status,
			Flags: []cli.Flag{
				stringFlag("url", "", "Add channel with `URL`"),
				stringFlag("service, s", "", "Add channel with `SERVICE`, user required"),
				stringFlag("username, u", "", "Add channel with `USER`, service required"),
			},
		},
	}
	app.Before = Init
	app.Run(os.Args)

}

func stringFlag(name, value, usage string) cli.StringFlag {
	return cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func intFlag(name string, value int, usage string) cli.IntFlag {
	return cli.IntFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}
