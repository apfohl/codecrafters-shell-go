package autocomplete

import (
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/app/file_system"
	"io"
	"iter"
	"os"
	"strings"
)

func Complete(
	prefix string,
	builtins map[string]func(iter.Seq[string], []string, io.WriteCloser, io.WriteCloser),
) string {
	if prefix == "" {
		return ""
	}

	//var suffixes []string
	suffixes := make(map[string]bool)
	for command := range builtins {
		after, found := strings.CutPrefix(command, prefix)
		if found {
			suffixes[after] = true
			//suffixes = append(suffixes, after)
		}
	}

	for _, executable := range file_system.FindExecutablesByPrefix(prefix) {
		after, found := strings.CutPrefix(executable, prefix)
		if found {
			suffixes[after] = true
			//suffixes = append(suffixes, after)
		}
	}

	if len(suffixes) == 1 {
		for suffix := range suffixes {
			return suffix
		}
		//return suffixes[0]
	}

	_, _ = fmt.Fprint(os.Stdout, "\a")

	return ""
}
