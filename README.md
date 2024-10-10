[![Go](https://github.com/RIDOS/manual_parse/actions/workflows/go.yml/badge.svg)](https://github.com/RIDOS/manual_parse/actions/workflows/go.yml)

## Manual parse
This is a test script written in Golang version 1.23. This script allows displaying the user's name via a flag.

## How it works
Build your app:
```bash
go build main.go -o manual_parse
```

Run your app:
```bash
./manual_parse --help
```

## How to test it

```bash
go test -v
```

> But also you can view precentations covired by this script.
>
> ```bash
> go test -coverprofile cover.out && go tool cover -html=cover.out
> ```

## Links
https://github.com/practicalgo/code
