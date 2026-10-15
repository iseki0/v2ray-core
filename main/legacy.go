package main

import (
	"flag"
	"fmt"
	"io"
	"strings"
)

// legacyArgs adapts the flag-only CLI used by Qv2ray to the v5 commands.
// Explicit subcommands and top-level help retain their upstream behavior.
func legacyArgs(args []string) ([]string, error) {
	if len(args) == 0 || !strings.HasPrefix(args[0], "-") || args[0] == "-h" || args[0] == "--help" || args[0] == "-help" {
		return args, nil
	}
	flags := flag.NewFlagSet("v2ray", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	version := flags.Bool("version", false, "print version")
	test := flags.Bool("test", false, "test configuration")
	forward := []string{}
	for _, name := range []string{"config", "c", "confdir", "d", "format"} {
		flags.Func(name, "configuration option", func(value string) error {
			forward = append(forward, "-"+name+"="+value)
			return nil
		})
	}
	recursive := flags.Bool("r", false, "load configuration directories recursively")
	if err := flags.Parse(args); err != nil {
		return nil, err
	}
	if flags.NArg() != 0 {
		return nil, fmt.Errorf("unexpected arguments: %v", flags.Args())
	}
	if *version {
		return []string{"version"}, nil
	}
	command := "run"
	if *test {
		command = "test"
	}
	if *recursive {
		forward = append(forward, "-r")
	}
	return append([]string{command}, forward...), nil
}
