package golang

func IsValid(s string) bool {

	const (
		PARENTHESIS_OPEN  = 40
		PARENTHESIS_CLOSE = 41
		BRACE_OPEN        = 123
		BRACE_CLOSE       = 125
		BRACKET_OPEN      = 91
		BRACKET_CLOSE     = 93
	)

	toClose := []rune{0}

	for _, e := range s {
		switch e {

		case PARENTHESIS_OPEN:
			toClose = append(toClose, e)

		case PARENTHESIS_CLOSE:
			if toClose[len(toClose)-1] != PARENTHESIS_OPEN {
				return false
			}
			toClose = toClose[:len(toClose)-1]

		case BRACE_OPEN:
			toClose = append(toClose, e)

		case BRACE_CLOSE:
			if toClose[len(toClose)-1] != BRACE_OPEN {
				return false
			}
			toClose = toClose[:len(toClose)-1]

		case BRACKET_OPEN:
			toClose = append(toClose, e)

		case BRACKET_CLOSE:
			if toClose[len(toClose)-1] != BRACKET_OPEN {
				return false
			}
			toClose = toClose[:len(toClose)-1]
		}
	}

	return len(toClose) == 1
}
