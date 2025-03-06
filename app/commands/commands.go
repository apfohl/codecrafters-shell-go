package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Exit(args []string) {
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

func Echo(args []string) {
	_, _ = fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))
}
