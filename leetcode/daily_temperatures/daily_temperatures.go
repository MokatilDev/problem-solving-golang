package daily_temperatures

func dailyTemperatures(temperatures []int) []int {
	if len(temperatures) == 0 {
		return []int{}
	}

	answer := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures))

	for i, temperature := range temperatures {
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[prev] = i - prev
		}

		stack = append(stack, i)
	}

	return answer
}
