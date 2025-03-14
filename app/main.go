package main

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"maps"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/arguments"
	"github.com/codecrafters-io/shell-starter-go/app/commands"
)

func main() {
	builtins := map[string]func(iter.Seq[string], []string, io.WriteCloser, io.WriteCloser){
		"exit": commands.Exit,
		"echo": commands.Echo,
		"type": commands.Type,
		"pwd":  commands.Pwd,
		"cd":   commands.Cd,
	}

	for {
		_, _ = fmt.Fprint(os.Stdout, "$ ")

		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.Trim(input, "\n")

		if len(input) == 0 {
			continue
		}

		parsedInput := arguments.ParseArgs(input)
		commandName := parsedInput[0]
		var args []string
		if len(parsedInput) > 1 {
			args = parsedInput[1:]
		}

		redirect, err := arguments.FindOutputRedirect(args)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "%s: output redirect failure: %s\n", commandName, err.Error())
			continue
		}

		command, ok := builtins[commandName]
		if ok {
			if redirect.IsRedirect {
				var file *os.File
				var flags = os.O_WRONLY | os.O_CREATE
				if redirect.Append {
					flags |= os.O_APPEND
				}

				file, err = os.OpenFile(redirect.Destination, flags, 0644)
				if err != nil {
					_, _ = fmt.Fprintf(os.Stdout, "%s: output open redirect destination: %s\n", commandName, err.Error())
				}

				switch redirect.Direction {
				case 1:
					command(maps.Keys(builtins), redirect.CommandArgs, file, os.Stderr)
				case 2:
					command(maps.Keys(builtins), redirect.CommandArgs, os.Stdout, file)
				}

				file.Close()
				continue
			}

			command(maps.Keys(builtins), redirect.CommandArgs, os.Stdout, os.Stderr)
			continue
		}

		output, err := commands.Execute(commandName, redirect)
		if err != nil {
			continue
		}

		if len(output) > 0 {
			_, _ = fmt.Fprintf(os.Stdout, "%s\n", strings.Join(output, "\n"))
		}
	}
}
