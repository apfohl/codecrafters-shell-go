package arguments

func ParseArgs(input string) []string {
	inSingleQuotes := false
	inDoubleQuotes := false
	buffer := make([]rune, 0)

	args := make([]string, 0)
	for _, char := range input {
		if char == '"' {
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

		buffer = append(buffer, char)
	}

	if len(buffer) > 0 {
		args = append(args, string(buffer))
	}

	return args
}
