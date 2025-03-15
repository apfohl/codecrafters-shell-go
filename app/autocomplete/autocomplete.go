package autocomplete

import (
	"io"
	"iter"
	"strings"
)

func Complete(
	input string,
	builtins map[string]func(iter.Seq[string], []string, io.WriteCloser, io.WriteCloser),
) string {
	if input == "" {
		return ""
	}

	var suffixes []string
	for command := range builtins {
		after, found := strings.CutPrefix(command, input)
		if found {
			suffixes = append(suffixes, after)
		}
	}

	if len(suffixes) == 1 {
		return suffixes[0]
	}

	return ""
}
