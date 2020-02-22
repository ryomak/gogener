package app

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/ryomak/gogener/recipe"
	"github.com/ryomak/gogener/templater"
	"github.com/urfave/cli/v2"
)

func List(c *cli.Context) error {
	for key, templates := range recipe.Templates() {
		fmt.Printf("%s \n", aurora.Green(key))
		for _, v := range templates {
			fmt.Printf("- %s \n", v)
		}
	}
	return nil
}

func Create(c *cli.Context) error {
	if c.Args().Len() != 1 {
		fmt.Printf("expect args 1 but %d\n", c.Args().Len())
	}
	appFunc, err := recipe.GetAppTmplFunc(c.Args().First())
	if err != nil {
		fmt.Printf("Oops... cannot find templates: %s\n", aurora.Red(c.Args().First()))
		return err
	}
	return createApp(appFunc)(c)
}

func createApp(appTmplFunc func(c *cli.Context) *templater.AppTemplate) func(c *cli.Context) error {
	out := os.Stdout
	return func(c *cli.Context) error {

		appName := c.String("app")
		if appName == "" {
			fmt.Fprint(out, "appName -> ")
			fmt.Scan(&appName)
		}
		mod := c.String("mod")
		if mod == "" {
			fmt.Fprint(out, "modName -> ")
			fmt.Scan(&mod)
		}
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		appRecipe := recipe.New(out, recipe.Params{
			CurrentDir: wd,
			AppName:    appName,
			ModName:    mod,
		})
		if ok := appRecipe.Do(appTmplFunc(c)); !ok {
			fmt.Fprintf(out, "\n\n%s ", aurora.Red("Create Failed: "+appName))
		} else {
			fmt.Fprintf(out, "\n\n%s \n%s ", aurora.Green("Success!!!"), aurora.Cyan(fmt.Sprint("created ", appName, "  🎉 🎉\n\n")))
		}
		return nil
	}
}
