package main

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"maps"
	"os"
	"os/exec"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
)

func main() {
	commandMap := map[string]func(iter.Seq[string], []string){
		"exit": commands.Exit,
		"echo": commands.Echo,
		"type": commands.Type,
	}

	for {
		_, _ = fmt.Fprint(os.Stdout, "$ ")

		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.Trim(input, "\n")

		if len(input) == 0 {
			continue
		}

		parts := strings.Split(input, " ")
		cmd := parts[0]
		var args []string
		if len(parts) > 1 {
			args = parts[1:]
		}

		command, ok := commandMap[cmd]
		if ok {
			command(maps.Keys(commandMap), args)
			continue
		}

		c := exec.Command(cmd, args...)

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
			_, _ = fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)
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
