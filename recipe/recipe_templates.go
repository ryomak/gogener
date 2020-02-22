package recipe

import (
	"errors"

	"github.com/ryomak/gogener/recipe/local"
	"github.com/ryomak/gogener/recipe/remote"
	"github.com/ryomak/gogener/templater"
	"github.com/urfave/cli/v2"
)

var remoteTemplates RecipeTemplates
var localTemplates RecipeTemplates

func init() {
	remoteTemplates = remote.NewTemplates()
	localTemplates = local.NewTemplates()
}

type RecipeTemplates interface {
	List() []string
	IsExist(string) bool
	GetFunc(string) func(*cli.Context) *templater.AppTemplate
}

func Templates() map[string][]string {

	return map[string][]string{
		"local":  localTemplates.List(),
		"reomte": remoteTemplates.List(),
	}
}

func GetAppTmplFunc(tmplName string) (func(c *cli.Context) *templater.AppTemplate, error) {

	if localTemplates.IsExist(tmplName) {
		return localTemplates.GetFunc(tmplName), nil
	}
	if remoteTemplates.IsExist(tmplName) {
		return remoteTemplates.GetFunc(tmplName), nil
	}
	return nil, errors.New("cannot find this template.")
}
