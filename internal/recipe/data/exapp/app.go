package exapp

import (
	"path/filepath"

	"github.com/ryomak/go-app-cli/internal/templater"
	"github.com/urfave/cli/v2"
)

// AppTmplFunc is getting template for app
func AppTmplFunc(c *cli.Context) *templater.AppTemplate {
	return &templater.AppTemplate{
		Name: "ex-app",
		Tmpls: []*templater.Templater{
			{
				Tmpl: `
				package main

				import (
					"fmt"
				)

				func main() {
					fmt.Println("hello {{.AppName}}")
				}
				`,
				Dir:  filepath.Join(""),
				Name: "main.go",
			},
		},
	}
}
