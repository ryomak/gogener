## How to add remote App
for example creating exapp

- **Delims is “[[ ]]”**
- we can use templateMethod
  - AppName
  - ModName
  - ToCamel
  - ToLowerCamel

### 1. create project
```
.
├── app_template.yaml
├── bg.txt
├── go.mod
└── src
    └── main.go
```


#### app_template.yaml

```md
name: "exapp"
bg-file-path: "bg.txt"
templates:
  - "src/main.go"
  - "go.mod"
```
#### go.mod
```
module [[.ModName]]

go 1.12
```

#### src/main.go

```go:main.go

package main
import (
    "fmt"
)

func main() {
    fmt.Println("hello [[.AppName]]")
}

```

### 2.release by github pages
1. put templates to github repository 
2. release by github pages

### 3. update gogener file
append remote url of app_template.yaml

#### interal/recipe/remote/templates.go

```go
var remoteMap = map[string]string{
	"go-deep-util-example": "https://ryomak.github.io/templates-for-gogener/go-deep-util/app_template.yaml",
  "grpc-vue-go-example":  "https://ryomak.github.io/templates-for-gogener/grpc-vue-example/app_template.yaml",
+ "exapp": "{github pages url)}/app_template.yaml",
}
```

### example
ex [gogener-templates](https://github.com/ryomak/templates-for-gogener)
