package main

import (
	"os"

	"github.com/ryomak/gogener/app"
)

func main() {
	app := app.New("gogener", "CLI tool that automatically creates apps using templates", "1.0.0")
	app.Run(os.Args)
}
