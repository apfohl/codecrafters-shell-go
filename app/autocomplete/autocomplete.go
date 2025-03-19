package autocomplete

import (
	"fmt"
	"io"
	"iter"
	"os"
	"slices"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/file_system"
)

var tabbedPrefix string

func Complete(
	prefix string,
	builtins map[string]func(iter.Seq[string], []string, io.WriteCloser, io.WriteCloser),
) string {
	if prefix == "" {
		return ""
	}

	suffixes := make(map[string]bool)
	for command := range builtins {
		after, found := strings.CutPrefix(command, prefix)
		if found {
			suffixes[after] = true
		}
	}

	for _, executable := range file_system.FindExecutablesByPrefix(prefix) {
		after, found := strings.CutPrefix(executable, prefix)
		if found {
			suffixes[after] = true
		}
	}

	commands := make([]string, 0)
	for suffix := range suffixes {
		commands = append(commands, prefix+suffix)
	}

	slices.Sort(commands)

	longestPrefix := findLongestPrefix(commands)

	if len(suffixes) > 1 {
		if len(tabbedPrefix) > 0 && tabbedPrefix == prefix {
			_, _ = fmt.Fprintf(os.Stdout, "\r\n%s\r\n$ %s", strings.Join(commands, "  "), prefix)
			return ""
		}

		_, _ = fmt.Fprint(os.Stdout, "\a")

		suffix, _ := strings.CutPrefix(longestPrefix, prefix)

		//suffix, _ := strings.CutPrefix(commands[0], prefix)
		tabbedPrefix = prefix + suffix
		return suffix
	}

	if len(suffixes) == 1 {
		tabbedPrefix = ""
		for suffix := range suffixes {
			return suffix + " "
		}
	}

	_, _ = fmt.Fprint(os.Stdout, "\a")
	return ""
}

func findLongestPrefix(commands []string) string {
	if len(commands) == 0 {
		return ""
	}

	prefix := ""

	for i, char := range commands[0] {
		for _, command := range commands[1:] {
			if []int32(command)[i] != char {
				return prefix
			}
		}
		prefix += string(char)
	}

	return prefix
}
