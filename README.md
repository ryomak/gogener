# gogener
[![Build Status](https://github.com/ryomak/gogener/workflows/Test/badge.svg)](https://github.com/ryomak/gogener/actions?workflow=Test)
[![Coverage Status](https://coveralls.io/repos/github/ryomak/gogener/badge.svg?branch=master)](https://coveralls.io/github/ryomak/gogener?branch=master)
[![GoDoc](https://godoc.org/github.com/ryomak/gogener?status.svg)](https://godoc.org/github.com/ryomak/gogener)
[![GoReport](https://goreportcard.com/badge/github.com/ryomak/gogener)](https://goreportcard.com/report/github.com/ryomak/gogener)

gogener is CLI tool that automatically creates apps using templates

![test](https://user-images.githubusercontent.com/21288308/70862524-d8b16200-1f80-11ea-812d-8ee6a92a140b.gif)

## install

```
$ go get github.com/ryomak/gogener/cmd/gogener
```

## usage

```
NAME:
   gogener - CLI tool that automatically creates apps using templates

USAGE:
   gogener [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
   create   create App with templates
   list     show template list
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## example

```
$ gogener create  -app example -mod github.com/ryomak/example ryomak-app
```


## how add App
https://ryomak.github.io/gogener


## WIP
- add test
- add other templates(now. only ryomak-app)

## License
MIT
