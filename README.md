# go-app-cli
[![Build Status](https://github.com/ryomak/go-app-cli/workflows/go/badge.svg)](https://github.com/ryomak/go-app-cli/actions?workflow=Test)
[![Coverage Status](https://coveralls.io/repos/github/ryomak/go-app-cli/badge.svg?branch=master)](https://coveralls.io/github/ryomak/go-app-cli?branch=master)
[![GoDoc](https://godoc.org/github.com/ryomak/go-app-cli?status.svg)](https://godoc.org/github.com/ryomak/go-app-cli)
[![GoReport](https://goreportcard.com/badge/github.com/ryomak/go-app-cli)](https://goreportcard.com/report/github.com/ryomak/go-app-cli)

go-app-cli is making project with template


## install

```
$ go get github.com/ryomak/go-app-cli/cmd/go-app-cli
```

## usage

```
NAME:
   go-app-cli create - create project

USAGE:
   go-app-cli create command [command options] [arguments...]

COMMANDS:
   ryomak-app  create a ryomak-app template
   help, h     Shows a list of commands or help for one command

OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## example

```
$ go-app-cli create  -app example -mod github.com/ryomak/example ryomak-app
```

## WIP
- add test
- add other templates(now. only ryomak-app)

## License
MIT
