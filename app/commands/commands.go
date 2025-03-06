package commands

import (
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"
)

func Exit(_ iter.Seq[string], args []string) {
	if len(args) == 0 {
		os.Exit(0)
	}

	code, err := strconv.Atoi(args[0])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s: could not be parsed\n", args[0])
		os.Exit(-1)
	}

	os.Exit(code)
}

func Echo(_ iter.Seq[string], args []string) {
	_, _ = fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))
}

func Type(commands iter.Seq[string], args []string) {
	if len(args) != 1 {
		_, _ = fmt.Fprintf(os.Stderr, "type: invalid number of arguments (%d of 1)\n", len(args))
		return
	}

	cmd := args[0]

	for command := range commands {
		if command == cmd {
			_, _ = fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", cmd)
			return
		}
	}

	pathEnv, found := os.LookupEnv("PATH")
	if !found {
		_, _ = fmt.Fprintf(os.Stdout, "env variable not found: PATH\n")
		return
	}

	for _, directory := range strings.Split(pathEnv, ":") {
		entries, err := os.ReadDir(directory)
		if err != nil {
			continue
			_, _ = fmt.Fprintf(os.Stderr, "could not read directory: %s\n", directory)
			os.Exit(-1)
		}

		for _, entry := range entries {
			if !entry.IsDir() && entry.Name() == cmd {
				_, _ = fmt.Fprintf(os.Stdout, "%s is %s\n", cmd, fmt.Sprintf("%s/%s", directory, cmd))
				return
			}
		}
	}

	_, _ = fmt.Fprintf(os.Stdout, "%s: not found\n", cmd)
}
