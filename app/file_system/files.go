package file_system

import (
	"fmt"
	"os"
	"strings"
)

func FindExecutable(name string) (string, error) {
	pathEnv, found := os.LookupEnv("PATH")
	if !found {
		_, _ = fmt.Fprintf(os.Stderr, "env variable not found: PATH\n")
		os.Exit(-1)
	}

	for directory := range strings.SplitSeq(pathEnv, ":") {
		entries, err := os.ReadDir(directory)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			if !entry.IsDir() && entry.Name() == name {
				return fmt.Sprintf("%s/%s", directory, name), nil
			}
		}
	}

	return "", fmt.Errorf("%s: not found", name)
}

func FindExecutablesByPrefix(prefix string) []string {
	pathEnv, found := os.LookupEnv("PATH")
	if !found {
		_, _ = fmt.Fprintf(os.Stderr, "env variable not found: PATH\n")
		os.Exit(-1)
	}

	var executables []string

	for directory := range strings.SplitSeq(pathEnv, ":") {
		entries, err := os.ReadDir(directory)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			if !entry.IsDir() && strings.HasPrefix(entry.Name(), prefix) {
				executables = append(executables, entry.Name())
			}
		}
	}

	return executables
}
