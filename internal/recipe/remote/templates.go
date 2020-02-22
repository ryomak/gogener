package remote

import (
	"github.com/ryomak/gogener/internal/templater"
	"github.com/urfave/cli/v2"
)

var remoteMap = map[string]string{
	"ryomak/go-deep-util-example": "https://ryomak.github.io/templates-for-gogener/go-deep-util/app_template.yaml",
	"ryomak/grpc-vue-go-example":  "https://ryomak.github.io/templates-for-gogener/grpc-vue-example/app_template.yaml",
	"ryomak/go-p2pchat":           "https://ryomak.github.io/templates-for-gogener/go-p2pchat/app_template.yaml",
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
