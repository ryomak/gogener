package ryomak

import (
	"path/filepath"

	"github.com/ryomak/gogener/internal/templater"
	"github.com/urfave/cli/v2"
)

// AppTmplFunc is getting template for app
func AppTmplFunc(c *cli.Context) *templater.AppTemplate {
	return &templater.AppTemplate{
		Name: "ryomak-app",
		Bg: "  _ __   __  __      ___     ___ ___       __    \\ \\ \\/\\     \n" +
			"/\\`'__\\/\\ \\/\\ \\    / __`\\ /' __` __`\\   /'__`\\    \\ \\ , <     \n" +
			"\\ \\ \\/ \\ \\ \\_\\ \\  /\\ \\L\\ \\/\\ \\/\\ \\/\\ \\ /\\ \\L\\.\\_   \\ \\ \\\\`\\   \n" +
			" \\ \\_\\  \\/`____ \\ \\ \\____/\\ \\_\\ \\_\\ \\_\\\\ \\__/.\\_\\   \\ \\_\\ \\_\\ \n" +
			"  \\/_/   `/___/> \\ \\/___/  \\/_/\\/_/\\/_/ \\/__/\\/_/    \\/_/\\/_/ \n" +
			"            /\\___/                                            \n" +
			"            \\/__/",
		Tmpls: []*templater.Templater{
			{
				Tmpl: localTmpl{modTempl},
				Dir:  filepath.Join(""),
				Name: "go.mod",
			},
			{
				Tmpl: localTmpl{mainTmpl},
				Dir:  filepath.Join("src"),
				Name: "main.go",
			},
			{
				Tmpl: localTmpl{envTmpl},
				Dir:  filepath.Join(""),
				Name: ".env",
			},
			{
				Tmpl: localTmpl{makeFileTmpl},
				Dir:  filepath.Join(""),
				Name: "Makefile",
			},
			{
				Tmpl: localTmpl{usecaseTmpl},
				Dir:  filepath.Join("src", "application", "usecase"),
				Name: "user_usecase.go",
			},
			{
				Tmpl: localTmpl{domainEntityUserTmpl},
				Dir:  filepath.Join("src", "domain", "user"),
				Name: "user.go",
			},
			{
				Tmpl: localTmpl{domainUserRepositoryTmpl},
				Dir:  filepath.Join("src", "domain", "user"),
				Name: "user_repository.go",
			},
			{
				Tmpl: localTmpl{domainUserServiceTmpl},
				Dir:  filepath.Join("src", "domain", "user"),
				Name: "user_service.go",
			},
			{
				Tmpl: localTmpl{infraUserRepositoryTmpl},
				Dir:  filepath.Join("src", "infrastructure", "repository"),
				Name: "user_repository.go",
			},
			{
				Tmpl: localTmpl{handlerTmpl},
				Dir:  filepath.Join("src", "interface", "handler"),
				Name: "handler.go",
			},
			{
				Tmpl: localTmpl{userHandlerTmpl},
				Dir:  filepath.Join("src", "interface", "handler"),
				Name: "user_handler.go",
			},
			{
				Tmpl: localTmpl{routerTmpl},
				Dir:  filepath.Join("src", "interface", "router"),
				Name: "router.go",
			},
			{
				Tmpl: localTmpl{loggerTmpl},
				Dir:  filepath.Join("src", "internal", "logger"),
				Name: "logger.go",
			},
			{
				Tmpl: localTmpl{middlewareTmpl},
				Dir:  filepath.Join("src", "internal", "middleware"),
				Name: "middleware.go",
			},
		},
	}
}

type localTmpl struct {
	tmplStr string
}

func (s localTmpl) String() (string, error) {
	// ceternly error nil because local str
	return s.tmplStr, nil
}
