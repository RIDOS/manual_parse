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
	fs.IntVar(&c.numTimes, "n", 0, "Number of times to greet")
	fs.StringVar(&c.filePath, "o", "", "Folder path for your HTML file")

	err := fs.Parse(args)
	if err != nil {
		return c, err
	}

	if fs.NArg() != 0 {
		return c, errPosArgSpecified
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

func runCmd(r io.Reader, w io.Writer, c config) error {
	name, err := getName(r, w)
	if err != nil {
		return err
	}
	greetUser(c, name, w)
	if c.filePath != "" {
		createTemplate(c, name)
	}
	return nil
}

func greetUser(c config, name string, w io.Writer) {
	msg := fmt.Sprintf("Nite to meet you %s\n", name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}

func createTemplate(c config, name string) {
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
	tmpl.Execute(file, name)
}

var errPosArgSpecified = errors.New("Positional arguments specified")

func main() {
	c, err := parseArgs(os.Stdout, os.Args[1:])
	if err != nil {
		if errors.Is(err, errPosArgSpecified) {
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
