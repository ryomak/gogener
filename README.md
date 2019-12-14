# go-cli
go-cli is making project with template


## install

```
$ GO111MODULE=on go get github.com/ryomak/go-cli
```

## usage

```
NAME:
   go-cli create - create project

USAGE:
   go-cli create command [command options] [arguments...]

COMMANDS:
   ryomak-app  create a ryomak-app template
   help, h     Shows a list of commands or help for one command

OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## example

```
$ go-cli create ryomak-app -app example -mod github.com/ryomak/example
```

## License
MIT
