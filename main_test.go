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
		arg      []string
		name     string
		output   string
		exitCode int
	}{
		{
			arg:      []string{"-n", "5"},
			name:     "a.imaev",
			output:   "Your name please? Press the Enter key when done.\n" + strings.Repeat("Nite to meet you a.imaev\n", 5),
			exitCode: 0,
		},
		{
			arg:      []string{"-n", "-5"},
			output:   "Must specify a number greater than 0\n",
			exitCode: 1,
		},
		{
			arg:      []string{"-h", ""},
			output:   "Usage of greeter:\n  -n int\n    \tNumber of times to greet\nflag: help requested\n",
			exitCode: 1,
		},
		{
			arg:      []string{"--help", ""},
			output:   "Usage of greeter:\n  -n int\n    \tNumber of times to greet\nflag: help requested\n",
			exitCode: 1,
		},
		{
			arg:      []string{"-h", "-n"},
			output:   "Usage of greeter:\n  -n int\n    \tNumber of times to greet\nflag: help requested\n",
			exitCode: 1,
		},
	}

	for _, tc := range test {
		cmd := exec.Command("./console-prog", tc.arg...)
		cmd.Stdin = strings.NewReader(tc.name)

		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()

		gotMsg := out.String()
		if gotMsg != tc.output {
			t.Errorf("Expected stdout message to b: \n%v, Got \n%v\n", tc.output, gotMsg)
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
