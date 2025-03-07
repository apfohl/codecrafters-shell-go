package arguments

func ParseArgs(input string) []string {
	inLiteral := false
	literal := make([]rune, 0)
	buffer := make([]rune, 0)
	firstSpaceAdded := false

	args := make([]string, 0)
	for _, s := range input {
		if s == '\'' && inLiteral {
			inLiteral = false
			args = append(args, string(literal))
			literal = make([]rune, 0)
			continue
		}

		if s == '\'' {
			if len(buffer) > 0 {
				args = append(args, string(buffer))
				buffer = make([]rune, 0)
			}

			inLiteral = true
			continue
		}

		if inLiteral {
			literal = append(literal, s)
			continue
		}

		if s == ' ' && !firstSpaceAdded {
			buffer = append(buffer, s)
			firstSpaceAdded = true
			continue
		}

		if s == ' ' && len(buffer) > 0 {
			args = append(args, string(buffer))
			buffer = make([]rune, 0)
			continue
		}

		if s == ' ' {
			continue
		}
		firstSpaceAdded = false

		buffer = append(buffer, s)
	}

	if len(buffer) > 0 {
		args = append(args, string(buffer))
	}

	return args
}
