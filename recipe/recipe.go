package recipe

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/logrusorgru/aurora"
	"github.com/ryomak/gogener/templater"
)

// Recipe is template execute interface like chef
type Recipe interface {
	Do(*templater.AppTemplate) bool
}

// Params is setting sturct for template
type Params struct {
	IsRemote   bool
	CurrentDir string
	AppName    string
	ModName    string
}

type appRecipe struct {
	out    io.Writer
	params Params
}

// New is creating Recipe
func New(out io.Writer, params Params) Recipe {
	return &appRecipe{
		out:    out,
		params: params,
	}
}

func (r *appRecipe) Do(appTmpl *templater.AppTemplate) bool {
	if appTmpl == nil {
		return false
	}
	fmt.Fprintf(r.out, "%s\n\n", aurora.Green(fmt.Sprint("Fetch template: ", appTmpl.Name)))
	if appTmpl.Bg != "" {
		fmt.Fprintf(r.out, "%s\n\n", appTmpl.Bg)
	}
	fmt.Fprintf(r.out, "%s...\n", aurora.Green(fmt.Sprint("Start Creating App: ", r.params.AppName)))

	absPath := filepath.Join(r.params.CurrentDir, r.params.AppName)

	ok := true
	for _, v := range appTmpl.Tmpls {
		if err := v.Create(absPath, r.params); err != nil {
			fmt.Fprintf(r.out, "[%s] %s \n ", aurora.Red("×"), filepath.Join(v.Dir, v.Name))
			ok = false
			fmt.Println(err.Error())
		} else {
			fmt.Fprintf(r.out, "[%s] %s \n", aurora.Green("✔︎"), filepath.Join(v.Dir, v.Name))
		}
	}
	return ok
}
