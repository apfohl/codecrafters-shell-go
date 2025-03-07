package main

import (
	"bufio"
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/app/arguments"
	"io"
	"iter"
	"maps"
	"os"
	"os/exec"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
)

func main() {
	builtins := map[string]func(iter.Seq[string], []string){
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

		command, ok := builtins[commandName]
		if ok {
			command(maps.Keys(builtins), args)
			continue
		}

		c := exec.Command(commandName, args...)

		var stdout io.ReadCloser
		stdout, err := c.StdoutPipe()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(-1)
		}

		done := make(chan []string)
		scanner := bufio.NewScanner(stdout)

		go func() {
			var lines []string

			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}

			done <- lines
		}()

		if err = c.Start(); err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "%s: command not found\n", commandName)
			continue
		}

		lines := <-done

		if err = c.Wait(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(-1)
		}

		_, _ = fmt.Fprintf(os.Stdout, "%s\n", strings.Join(lines, "\n"))
	}
}
