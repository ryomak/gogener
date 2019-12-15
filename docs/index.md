## How to add project
ex)create "ex-app"
1. make directory```internal/recipe/data/exapp```
2. cd exapp && touch app.go

```go:app.go
func AppTmplFunc(c *cli.Context) *templater.AppTemplate {
	return &templater.AppTemplate{
		Name: "ex-app",
		Tmpls: []*templater.Templater{
			{
                Tmpl: `
                package main
                import (
                    "fmt"
                )

                func init() {
                    gotenv.Load()
                }

                func main() {
                    fmt.Println("hello {{.AppName}}")
                }
                `,
				Dir:  filepath.Join(""),
				Name: "main.go",
			},
		},
	}
}
```

- template can use ".AppName" and ".ModName" default .
- Dir is Relative path from ".AppName"
- name is file name.

3. change ```interal/app/app.go```
- add new subcommand

```go:app.go
func commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "create",
			Usage: "create project",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "mod",
					Value: "",
					Usage: "mod use",
				},
				&cli.StringFlag{
					Name:  "app",
					Value: "",
					Usage: "app name",
				},
			},
			Subcommands: []*cli.Command{
				{
					Name:   "ryomak-app",
					Usage:  "create a ryomak-app template",
					Action: createApp(ryomak.AppTmplFunc),
				},
                +{
                +    Name:   "ex-app",
				+	Usage:  "ex a new-app template",
				+	Action: createApp(exapp.AppTmplFunc),
                +},
			},
		},
```