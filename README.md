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

## Commands

### Flag h

If you run app with flag `-h` or `--help`:
```bash
A greeter application whitch prints the name you entered a specified number of times.

Usage of greeter: <option> [name]

Options: 
  -n int
        Number of times to greet
  -o string
        Folder path for your HTML file
```

### Flag n

Flag `-n` can write on console "Nice to meet you `<user_name>`. Example:
```bash
> ./console-prog -n 3 "Richard"
Nite to meet you Richard
Nite to meet you Richard
Nite to meet you Richard
```

### Flag o

Flag `-o` can create `html` file. Example:
```bash
➜  manual_parse git:(main) ✗ ./main -o ./
Your name please? Press the Enter key when done.
Richard
```
In output - create file with:
```html
<h1>Hello Richard!</h1>
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
