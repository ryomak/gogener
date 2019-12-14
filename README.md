# go-app-cli
go-app-cli is making project with template


## install

```
$ GO111MODULE=on go get github.com/ryomak/go-app-cli/cmd/go-app-cli
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
