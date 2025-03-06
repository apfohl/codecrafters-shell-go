package main

import (
	"bufio"
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/app/commands"
	"os"
	"strings"
)

func main() {
	commandMap := map[string]func([]string){
		"exit": commands.Exit,
		"echo": commands.Echo,
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
			os.Exit(-1)
		}

		command(args)
	}
}
