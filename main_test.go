package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

// info: Build this before start
func TestMain(t *testing.T) {
	test := []struct {
		arg      string
		name     string
		output   string
		exitCode int
	}{
		{
			arg:      "5",
			name:     "a.imaev",
			output:   "Your name please? Press the Enter key when done.\n" + strings.Repeat("Nite to meet you a.imaev\n", 5),
			exitCode: 0,
		},
		{
			arg:      "-5",
			output:   "Must specify a number greater than 0\nUsage: ./main <integer> [-h|--help]\nA greeter application which prints the name you entered\n<integer> number of times.\n",
			exitCode: 1,
		},
		{
			arg:      "-h",
			output:   "Must specify a number greater than 0\nUsage: ./main <integer> [-h|--help]\nA greeter application which prints the name you entered\n<integer> number of times.\n",
			exitCode: 1,
		},
		{
			arg:      "--help",
			output:   "Must specify a number greater than 0\nUsage: ./main <integer> [-h|--help]\nA greeter application which prints the name you entered\n<integer> number of times.\n",
			exitCode: 1,
		},
	}

	for _, tc := range test {
		cmd := exec.Command("./console-prog", tc.arg)
		cmd.Stdin = strings.NewReader(tc.name)

		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()

		gotMsg := out.String()
		if gotMsg != tc.output {
			t.Errorf("Expected stdout message to b: \n%v, Got \n%v\n", gotMsg, tc.output)
		}
		out.Reset()

		if err != nil {
			gotExitCode := err.(*exec.ExitError).ExitCode()
			if tc.exitCode != gotExitCode {
				t.Fatalf("Expected exception to b: %v, Got %v\n", tc.exitCode, err.Error())
			}
		}
	}
}
