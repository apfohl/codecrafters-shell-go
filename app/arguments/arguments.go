package arguments

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
