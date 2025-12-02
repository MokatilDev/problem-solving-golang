package climbingStairs

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	step1 := 1
	step2 := 2
	current := 0

	for i := 3; i <= 3; i++ {
		current = step1 + step2

		step1 = step2
		step2 = current
	}

	return current
}
