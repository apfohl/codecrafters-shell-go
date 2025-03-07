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

func Cd(_ iter.Seq[string], args []string) {
	if len(args) == 0 {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "cd: user home directory not found\n")
			return
		}

		err = os.Chdir(homeDir)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", homeDir)
			return
		}
	}

	if len(args) > 1 {
		_, _ = fmt.Fprintf(os.Stderr, "cd: invalid number of arguments (%d of 1)\n", len(args))
		return
	}

	segments := strings.Split(args[0], "/")
	path, err := initializePath(segments)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "cd: %s:could not parse input\n", args[0])
		return
	}

	if segments[0] == "" {
		segments = segments[1:]
	}

	for _, segment := range segments {
		switch segment {
		case "..":
			if len(path) > 0 {
				path = path[:len(path)-1]
			}
		case ".":
			continue
		default:
			path = append(path, segment)
		}
	}

	directory := strings.Join(path, "/")
	directory = fmt.Sprintf("/%s", directory)

	err = os.Chdir(directory)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", directory)
		return
	}
}

func initializePath(segments []string) ([]string, error) {
	if segments[0] == "" {
		return make([]string, 0), nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	path := strings.Split(wd, "/")
	return path[1:], nil
}
