package ryomak

import (
	"path/filepath"

	"github.com/ryomak/go-cli/internal/templater"
	"github.com/urfave/cli/v2"
)

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
				Tmpl: modTempl,
				Dir:  filepath.Join(""),
				Name: "go.mod",
			},
			{
				Tmpl: mainTmpl,
				Dir:  filepath.Join("src"),
				Name: "main.go",
			},
			{
				Tmpl: envTmpl,
				Dir:  filepath.Join(""),
				Name: ".env",
			},
			{
				Tmpl: makeFileTmpl,
				Dir:  filepath.Join(""),
				Name: "Makefile",
			},
			{
				Tmpl: usecaseTmpl,
				Dir:  filepath.Join("src", "application", "usecase"),
				Name: "user_usecase.go",
			},
			{
				Tmpl: domainEntityUserTmpl,
				Dir:  filepath.Join("src", "domain", "user"),
				Name: "user.go",
			},
			{
				Tmpl: domainUserRepositoryTmpl,
				Dir:  filepath.Join("src", "domain", "user"),
				Name: "user_repository.go",
			},
			{
				Tmpl: domainUserServiceTmpl,
				Dir:  filepath.Join("src", "domain", "user"),
				Name: "user_service.go",
			},
			{
				Tmpl: infraUserRepositoryTmpl,
				Dir:  filepath.Join("src", "infrastructure", "repository"),
				Name: "user_repository.go",
			},
			{
				Tmpl: handlerTmpl,
				Dir:  filepath.Join("src", "interface", "handler"),
				Name: "handler.go",
			},
			{
				Tmpl: userHandlerTmpl,
				Dir:  filepath.Join("src", "interface", "handler"),
				Name: "user_handler.go",
			},
			{
				Tmpl: routerTmpl,
				Dir:  filepath.Join("src", "interface", "router"),
				Name: "router.go",
			},
			{
				Tmpl: loggerTmpl,
				Dir:  filepath.Join("src", "internal", "logger"),
				Name: "logger.go",
			},
			{
				Tmpl: middlewareTmpl,
				Dir:  filepath.Join("src", "internal", "middleware"),
				Name: "middleware.go",
			},
		},
	}
}
