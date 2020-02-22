package local

import (
	"github.com/ryomak/gogener/recipe/local/ryomak"
	"github.com/ryomak/gogener/templater"
	"github.com/urfave/cli/v2"
)

var localMap = map[string]func(c *cli.Context) *templater.AppTemplate{
	"ryomak-app": ryomak.AppTmplFunc,
}

type templates struct{}

func NewTemplates() templates {
	return templates{}
}

func (templates) List() []string {
	remotes := make([]string, 0, len(localMap))
	for k := range localMap {
		remotes = append(remotes, k)
	}
	return remotes
}

func (templates) IsExist(key string) bool {
	_, ok := localMap[key]
	return ok
}

func (templates) GetFunc(key string) func(c *cli.Context) *templater.AppTemplate {
	f, ok := localMap[key]
	if !ok {
		return nil
	}
	return f
}
