package counting_bits

func countBits(n int) []int {
	res := make([]int, n+1)

	for i := 0; i <= n; i++ {
		count := 0
		num := i

		for num > 0 {
			count += num & 1
			num >>= 1
		}

		res[i] = count
	}

	return res
}
