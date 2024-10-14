package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
)

type config struct {
	numTimes int
	name     string
	filePath string
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) && c.filePath == "" {
		return errors.New("Must specify a number greater than 0")
	}
	return nil
}

func parseArgs(w io.Writer, args []string) (config, error) {
	c := config{}
	fs := flag.NewFlagSet("greeter", flag.ContinueOnError)

	fs.SetOutput(w)
	fs.Usage = func() {
		var usgaeString = `
A greeter application whitch prints the name you entered a specified number of times.

Usage of %s: <option> [name]
`
		fmt.Fprintf(w, usgaeString, fs.Name())
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options: ")
		fs.PrintDefaults()
	}
	fs.IntVar(&c.numTimes, "n", 0, "Number of times to greet")
	fs.StringVar(&c.filePath, "o", "", "Folder path for your HTML file")

	err := fs.Parse(args)
	if err != nil {
		return c, err
	}

	if fs.NArg() > 1 {
		return c, errInvalidPosArgSpeccified
	}
	if fs.NArg() == 1 {
		c.name = fs.Arg(0)
	}

	return c, nil
}

func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Your name please? Press the Enter key when done.\n"
	fmt.Fprintf(w, msg)

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("You didn't enter your name")
	}
	return name, nil
}

func runCmd(rd io.Reader, w io.Writer, c config) error {
	var err error
	if len(c.name) == 0 {
		c.name, err = getName(rd, w)
		if err != nil {
			return err
		}
	}
	if len(c.filePath) != 0 {
		createTemplate(c)
	}
	greetUser(c, w)
	return nil
}

func greetUser(c config, w io.Writer) {
	msg := fmt.Sprintf("Nite to meet you %s\n", c.name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}

func createTemplate(c config) {
	file, err := os.Create(c.filePath + "index.html")
	if err != nil {
		fmt.Printf("Can`t create file from path: %v\n", c.filePath)
		os.Exit(1)
	}
	tmpl, err := template.New("name").Parse("<h1>Hello {{ .}}!</h1>")
	if err != nil {
		fmt.Printf("Can`t implode name in file.")
		os.Exit(1)
	}
	tmpl.Execute(file, c.name)
}

var errInvalidPosArgSpeccified = errors.New("More than one positional argument specified")

func main() {
	c, err := parseArgs(os.Stdout, os.Args[1:])
	if err != nil {
		if errors.Is(err, errInvalidPosArgSpeccified) {
			fmt.Fprintln(os.Stdout, err)
		}
		os.Exit(1)
	}
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
