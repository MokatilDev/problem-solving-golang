package pivot

func pivotInteger(n int) int {
	for pivot := 1; pivot <= n; pivot++ {
		sum1 := 0
		sum2 := 0
		for i := 1; i <= pivot; i++ {
			sum1 += i
		}
		for j := pivot; j <= n; j++ {
			sum2 += j
		}

		if sum1 == sum2 {
			return pivot
		}
	}

	return -1
}
