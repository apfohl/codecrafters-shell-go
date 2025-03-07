package arguments

import (
	"errors"
	"strconv"
	"strings"
)

func ParseArgs(input string) []string {
	inSingleQuotes := false
	inDoubleQuotes := false
	escaped := false
	buffer := make([]rune, 0)

	args := make([]string, 0)
	for _, char := range input {
		if escaped {
			if inSingleQuotes {
				buffer = append(buffer, '\\')
				buffer = append(buffer, char)
				escaped = false
				continue
			}

			if inDoubleQuotes {
				if char == '\\' || char == '"' || char == '$' {
					buffer = append(buffer, char)
					escaped = false
					continue
				}

				buffer = append(buffer, '\\')
				buffer = append(buffer, char)
				escaped = false
				continue
			}

			buffer = append(buffer, char)
			escaped = false
			continue
		}

		if char == '"' {
			if inSingleQuotes {
				buffer = append(buffer, char)
				continue
			}

			if inDoubleQuotes {
				inDoubleQuotes = false
				continue
			}

			inDoubleQuotes = true
			continue
		}

		if char == '\'' {
			if inDoubleQuotes {
				buffer = append(buffer, char)
				continue
			}

			if inSingleQuotes {
				inSingleQuotes = false
				continue
			}

			inSingleQuotes = true
			continue
		}

		if char == ' ' {
			if inDoubleQuotes || inSingleQuotes {
				buffer = append(buffer, char)
				continue
			}

			if len(buffer) > 0 {
				args = append(args, string(buffer))
				buffer = make([]rune, 0)
			}

			continue
		}

		if char == '\\' {
			escaped = true
			continue
		}

		buffer = append(buffer, char)
	}

	if len(buffer) > 0 {
		args = append(args, string(buffer))
	}

	return args
}

type Redirect struct {
	IsRedirect  bool
	CommandArgs []string
	Direction   int
	Destination string
}

func FindOutputRedirect(args []string) (Redirect, error) {
	for i, arg := range args {
		if strings.HasSuffix(arg, ">") {
			if i+1 > len(args)-1 {
				return Redirect{IsRedirect: true}, errors.New("no output redirect Destination given")
			}

			direction := 1
			if len(arg) == 2 {
				var err error
				direction, err = strconv.Atoi(string(arg[0]))
				if err != nil {
					return Redirect{IsRedirect: true}, err
				}
			}

			return Redirect{
				IsRedirect:  true,
				CommandArgs: args[:i],
				Direction:   direction,
				Destination: args[i+1],
			}, nil
		}
	}

	return Redirect{CommandArgs: args}, nil
}
