package brackets

func Bracket(s string) bool {
	stack := []rune{}

	for _, r := range s {
		switch r {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		case ')', ']', '}':
			l := len(stack)
			if l == 0 || stack[l-1] != r {
				return false
			}
			stack = stack[:l-1]
		}
	}

	return len(stack) == 0
}
