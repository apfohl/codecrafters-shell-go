package arguments

func ParseArgs(input string) []string {
	inLiteral := false
	buffer := make([]rune, 0)

	args := make([]string, 0)
	for _, char := range input {
		if char == '\'' {
			if inLiteral {
				inLiteral = false
				continue
			}

			inLiteral = true
			continue
		}

		if char == ' ' {
			if inLiteral {
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
