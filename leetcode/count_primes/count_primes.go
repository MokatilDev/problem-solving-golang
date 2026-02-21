package countp_rimes

func countPrimes(n int) int {
	if n <= 1 {
		return 0
	}

	isComposite := make([]bool, n)

	for i := 2; i*i <= n; i++ {
		if !isComposite[i] {
			for j := i * i; j <= n; j += i {
				isComposite[i] = true
			}
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if !isComposite[i] {
			count++
		}
	}

	return count
}
