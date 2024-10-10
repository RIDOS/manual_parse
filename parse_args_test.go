package main

import (
	"errors"
	"os"
	"testing"
)

type testConfig struct {
	args []string
	err  error
	config
}

func TestParseArgs(t *testing.T) {
	tests := []testConfig{
		{
			args:   []string{"-h"},
			err:    errors.New("flag: help requested"),
			config: config{numTimes: 0},
		},
		{
			args:   []string{"-n", "10"},
			err:    nil,
			config: config{numTimes: 10},
		},
		{
			args:   []string{"-n", "abc"},
			err:    errors.New("invalid value \"abc\" for flag -n: parse error"),
			config: config{numTimes: 0},
		},
		{
			args:   []string{"-n 1", "foo"},
			err:    errors.New("flag provided but not defined: -n 1"),
			config: config{numTimes: 0},
		},
	}

	for _, tc := range tests {
		c, err := parseArgs(os.Stdout, tc.args)
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error to be: \n%v, got: \n%v\n", tc.err, err)
		}
		if tc.err == nil && err != nil {
			t.Errorf("Expected niul error, got: %v\n", err)
		}
		if c.numTimes != tc.numTimes {
			t.Errorf("Expected numTimes to be: %v, got: %v\n", tc.numTimes, c.numTimes)
		}
	}
}
