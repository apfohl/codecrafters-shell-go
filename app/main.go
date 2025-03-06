package main

import (
	"bufio"
	"fmt"
	"iter"
	"maps"
	"os"
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

		parts := strings.Split(input, " ")
		cmd := parts[0]
		var args []string
		if len(parts) > 1 {
			args = parts[1:]
		}

		command, ok := commandMap[cmd]
		if !ok {
			_, _ = fmt.Fprintf(os.Stdout, "%s: command not found\n", input)
			continue
		}

		keys := maps.Keys(commandMap)

		command(keys, args)
	}
}
