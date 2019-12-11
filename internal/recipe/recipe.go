package recipe

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/logrusorgru/aurora"
	"github.com/ryomak/go-cli/internal/templater"
)

type Recipe interface {
	Do(*templater.AppTemplate) bool
}

type Params struct {
	CurrentDir string
	AppName    string
	ModName    string
}

type appRecipe struct {
	out    io.Writer
	params *Params
}

func New(out io.Writer, params *Params) Recipe {
	return &appRecipe{
		out:    out,
		params: params,
	}
}

func (r *appRecipe) Do(appTmpl *templater.AppTemplate) bool {
	fmt.Fprintf(r.out, "%s\n\n", aurora.Green(fmt.Sprint("Start Creating With Template:", appTmpl.Name)))
	if appTmpl.Bg != "" {
		fmt.Fprintf(r.out, "%s\n\n", appTmpl.Bg)
	}
	fmt.Fprintf(r.out, "%s...\n", aurora.Green(fmt.Sprint("Start Creating App:", r.params.AppName)))
	ok := true
	absPath := filepath.Join(r.params.CurrentDir, r.params.AppName)
	for _, v := range appTmpl.Tmpls {
		if err := v.Create(absPath, r.params); err != nil {
			fmt.Fprintf(r.out, "[%s] %s \n ", aurora.Red("×"), filepath.Join(v.Dir, v.Name))
			ok = false
		} else {
			fmt.Fprintf(r.out, "[%s] %s \n", aurora.Green("✔︎"), filepath.Join(v.Dir, v.Name))
		}
	}
	return ok
}
