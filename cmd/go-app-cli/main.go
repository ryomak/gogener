package main

import (
	"os"

	"github.com/ryomak/go-app-cli/internal/cmd"
)

func main() {
	app := cmd.New("go-app-cli", "create app project", "1.0.0")
	app.Run(os.Args)
}
