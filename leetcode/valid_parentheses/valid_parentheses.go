package valid_parentheses

func isValid(s string) bool {
	var stack []rune

	if len(s) < 2 {
		return false
	}

	for _, current := range s {
		if (current == '(') || (current == '[') || (current == '{') {
			stack = append(stack, current)
		} else if current == ')' || current == ']' || current == '}' {
			if len(stack) < 1 {
				return false
			}

			top := stack[len(stack)-1]

			if (current == ']' && top != '[') ||
				(current == '}' && top != '{') ||
				(current == ')' && top != '(') {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
