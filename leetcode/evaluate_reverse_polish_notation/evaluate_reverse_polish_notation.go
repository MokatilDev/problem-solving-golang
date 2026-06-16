package evaluate_reverse_polish_notation

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":

			if len(stack) < 2 {
				return 0
			}

			a := stack[len(stack)-2]
			b := stack[len(stack)-1]

			stack = stack[:len(stack)-2]
			result := 0

			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				result = a / b
			}

			stack = append(stack, result)

		default:
			number, _ := strconv.Atoi(token)
			stack = append(stack, number)
		}
	}

	return stack[0]
}
