package climbing_stairs

func climbStairs(n int) int {
	if n <= 2 && n >= 0 {
		return n
	}

	arr := make([]int, n+1)
	arr[1], arr[2] = 1, 2

	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n]
}

/* func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	step1 := 1
	step2 := 2
	current := 0

	for i := 3; i <= n; i++ {
		current = step1 + step2

		step1 = step2
		step2 = current
	}

	return current
} */
