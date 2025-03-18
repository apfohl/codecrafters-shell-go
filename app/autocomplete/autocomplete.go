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

	if len(suffixes) > 1 {
		if len(tabbedPrefix) > 0 && tabbedPrefix == prefix {
			printCompletions(prefix, suffixes)
			return ""
		}

		_, _ = fmt.Fprint(os.Stdout, "\a")
		tabbedPrefix = prefix
		return ""
	}

	if len(suffixes) == 1 {
		tabbedPrefix = ""
		for suffix := range suffixes {
			return suffix
		}
	}

	_, _ = fmt.Fprint(os.Stdout, "\a")
	return ""
}

func printCompletions(prefix string, suffixes map[string]bool) {
	_, _ = fmt.Fprint(os.Stdout, "\r\n")

	var buffer []string
	for suffix := range suffixes {
		buffer = append(buffer, prefix+suffix)
	}

	slices.Sort(buffer)

	_, _ = fmt.Fprintf(os.Stdout, "%s\r\n$ %s", strings.Join(buffer, "  "), prefix)
}
