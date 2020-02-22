## How to add local App
ex)create "ex-app"
1. make directory```internal/recipe/local/exapp```
2. cd exapp && touch app.go

```go:app.go
package exapp

import (
	"path/filepath"

	"github.com/ryomak/gogener/internal/templater"
	"github.com/urfave/cli/v2"
)

func AppTmplFunc(c *cli.Context) *templater.AppTemplate {
	return &templater.AppTemplate{
		Name: "ex-app",
		Tmpls: []*templater.Templater{
			{sv
                Tmpl: `
                package main
                import (
                    "fmt"
                )

                func main() {
                    fmt.Println("hello [[.AppName]]")
                }
                `,
				Dir:  filepath.Join(""),
				Name: "main.go",
			},
		},
	}
}
```
- **Delims is "[[ ]]"**
- template can use ".AppName" and ".ModName" default .
- Dir is Relative path from ".AppName"
- Name is file name.
- other templateMethod : toCamel, toLowerCamel (ex. [[.AppName | ToLoewerCamel]])

3. change ```interal/recipe/local/templates.go```
- add new AppTmplFunc

```go:templates.go
import(
	..
	+ "github.com/ryomak/gogener/internal/recipe/local/exapp"
)
var localMap = map[string]func(c *cli.Context) *templater.AppTemplate{
	"ryomak-app": ryomak.AppTmplFunc,
	+ "exapp": exapp.AppTmplFunc
}

```

4. ``` $ gogener create exapp ``` 