package main

import (
	"os"

	"github.com/ryomak/gogener/app"
)

var (
  Version = "unset"
)

func main() {
	app := app.New("gogener", "CLI tool that automatically creates apps using templates", Version)
	app.Run(os.Args)
}
