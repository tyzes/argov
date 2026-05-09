# argov - CLI flag parser

## Installation

```sh
go get github.com/tyzes/argov
```

## Examples
### General Usage

```go
str := argov.String([]string{"s", "string"}, "default")
num := argov.Int([]string{"i", "int"}, 1234)
slice := argov.StringSlice([]string{"slice"})

positionals, err := argov.Parse(os.Args[1:])
if err != nil {
	fmt.Println(err)
	os.Exit(1)
}
fmt.Println(*str, *num, *slice)
fmt.Println(positionals)
```

### Custom Parser
```go
p := argov.NewParser()
str := p.String([]string{"s", "string"}, "default")
positionals, err := p.Parse(os.Args[1:])
```

### Options and help menu generation

```go
str := argov.String([]string{"str"}, "", argov.Description("a string"), argov.Placeholder("STRING"))
slice := argov.StringSlice([]string{"slice"}, argov.Description("a slice"), argov.Required(), argov.SplitOn(',', ';'))
positionals, err := argov.Parse(os.Args[1:], argov.NoMixing())

if err != nil {
	fmt.Println("error:", err)
	fmt.Println("usage: programfile <options> arguments\n\nOPTIONS")
	fmt.Println(argov.GenerateHelp())
}
```

## Full Documentation
Full API documentation is available on [pkg.go.dev](https://pkg.go.dev/github.com/tyzes/argov).
