package main

import (
	"os"

	"github.com/ryomak/go-cli/internal/cmd"
)

func main() {
	app := cmd.New("go-cli", "create app project", "1.0.0")
	app.Run(os.Args)
}
