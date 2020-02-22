package app

import (
	"github.com/urfave/cli/v2"
)

// New is creating cli App
func New(name, usage, version string) *cli.App {
	app := cli.NewApp()
	app.Name = name
	app.Usage = usage
	app.Version = version
	app.Commands = commands()
	return app
}

func commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "create",
			Usage: "create App with templates",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "mod",
					Value: "",
					Usage: "mod use",
				},
				&cli.StringFlag{
					Name:  "app",
					Value: "",
					Usage: "app name",
				},
			},
			Action: Create,
		},
		{
			Name:   "list",
			Usage:  "show template list",
			Action: List,
		},
	}
}
