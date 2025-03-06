package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.Trim(input, "\n")

		switch input {
		case "exit 0":
			os.Exit(0)
		default:
			_, _ = fmt.Fprintf(os.Stdout, "%s: command not found\n", input)
		}
	}
}
