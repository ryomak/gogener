package app

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/ryomak/go-app-cli/internal/recipe"
	"github.com/ryomak/go-app-cli/internal/templater"
	"github.com/urfave/cli/v2"
)

func createApp(appTmplFunc func(c *cli.Context) *templater.AppTemplate) func(c *cli.Context) error {
	out := os.Stdout
	return func(c *cli.Context) error {

		appName := c.String("app")
		if appName == "" {
			fmt.Fprint(out, "appName ~>")
			fmt.Scan(&appName)
		}
		mod := c.String("mod")
		if mod == "" {
			fmt.Fprint(out, "modName ~>")
			fmt.Scan(&mod)
		}
		d, err := os.Getwd()
		if err != nil {
			return err
		}
		appRecipe := recipe.New(out, recipe.Params{
			CurrentDir: d,
			AppName:    appName,
			ModName:    mod,
		})
		if ok := appRecipe.Do(appTmplFunc(c)); !ok {
			fmt.Fprintf(out, "\n\n%s \n%s", aurora.Red("Create Faild"), aurora.Cyan(fmt.Sprint("created ", appName)))
		}
		fmt.Fprintf(out, "\n\n%s \n%s ", aurora.Green("Success!!!"), aurora.Cyan(fmt.Sprint("created ", appName, "  🎉 🎉\n\n")))
		return nil
	}
}
