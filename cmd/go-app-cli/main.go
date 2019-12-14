package main

import (
	"os"

	"github.com/ryomak/go-app-cli/internal/app"
)

func main() {
	app := app.New("go-app-cli", "create app project", "1.0.0")
	app.Run(os.Args)
}
