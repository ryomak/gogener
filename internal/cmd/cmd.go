package cmd

import (
	"github.com/ryomak/go-app-cli/internal/recipe/data/ryomak"
	"github.com/urfave/cli/v2"
)

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
			Usage: "create project",
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
			Subcommands: []*cli.Command{
				{
					Name:   "ryomak-app",
					Usage:  "create a ryomak-app template",
					Action: createApp(ryomak.AppTmplFunc),
				},
			},
		},
	}
}
