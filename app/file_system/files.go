package file_system

import (
	"errors"
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

	for _, directory := range strings.Split(pathEnv, ":") {
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

	return "", errors.New(fmt.Sprintf("%s: not found", name))
}
