package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return
	}

	input = strings.Trim(input, "\n")
	switch input {
	default:
		fmt.Fprintf(os.Stdout, "%s: command not found\n", input)
	}
}
