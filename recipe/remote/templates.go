package remote

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/ryomak/gogener/templater"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

var remoteMap = map[string]string{}

func init() {
	url := os.Getenv("GOGENER_REMOTE_MAP_URL")
	if url == "" {
		url = "https://ryomak.github.io/gogener-templates/map.yaml"
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Oops GOGENER_REMOTE_MAP_URL is invalid. detail: %s", err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Oops cannot read body. detail: %s", err.Error())
		os.Exit(1)
	}
	err = yaml.Unmarshal(data, &remoteMap)
	if err != nil {
		fmt.Printf("Oops cannot parse remote yaml. detail: %s", err.Error())
		os.Exit(1)
	}

}

type templates struct{}

func NewTemplates() templates {
	return templates{}
}

func (templates) List() []string {
	remotes := make([]string, 0, len(remoteMap))
	for k := range remoteMap {
		remotes = append(remotes, k)
	}
	return remotes
}

func (templates) IsExist(key string) bool {
	_, ok := remoteMap[key]
	return ok
}

func (templates) GetFunc(key string) func(c *cli.Context) *templater.AppTemplate {
	u, ok := remoteMap[key]
	if !ok {
		return nil
	}
	return RemoteAppTmplFunc(u)
}
