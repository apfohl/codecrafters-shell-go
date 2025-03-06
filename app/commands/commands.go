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

	_, _ = fmt.Fprintf(os.Stdout, "%s: not found\n", cmd)
}
