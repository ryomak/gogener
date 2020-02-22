package remote

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/ryomak/gogener/internal/templater"
	"github.com/urfave/cli/v2"
)

func RemoteAppTmplFunc(url string) func(c *cli.Context) *templater.AppTemplate {
	return func(c *cli.Context) *templater.AppTemplate {
		if url == "" {
			fmt.Print("app URL ~>")
			fmt.Scan(&url)
		}
		tClient, err := templater.NewRemoteClient(url)
		if err != nil {
			fmt.Printf("%s...\n", aurora.Red(fmt.Sprint("cannot parse url ", url)))
			os.Exit(1)
		}
		appTemplate, err := tClient.RemoteSettingToAppTemplate()
		if err != nil {
			fmt.Printf("Oops...cannot fetch %s/%s\n", aurora.Red(tClient.BaseURL), aurora.Red(tClient.SettingYmlName))
			fmt.Printf(err.Error())
			return nil
		}
		return appTemplate
	}
}
