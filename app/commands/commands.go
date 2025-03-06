package commands

import (
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/file_system"
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

	command, err := file_system.FindExecutable(cmd)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return
	}

	_, _ = fmt.Fprintf(os.Stdout, "%s is %s\n", cmd, command)
}

func Pwd(_ iter.Seq[string], _ []string) {
	directory, err := os.Getwd()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "can not read working directory\n")
		os.Exit(-1)
	}

	_, _ = fmt.Fprintf(os.Stdout, "%s\n", directory)
}
