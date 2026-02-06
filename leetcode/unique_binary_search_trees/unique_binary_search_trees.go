package unique_binary_search_trees

func numTrees(n int) int {
	numTree := make([]int, n+1)

	numTree[0] = 1
	if n >= 1 {
		numTree[1] = 1
	}

	for i := 2; i <= n; i++ {
		total := 0

		for j := 1; j <= i; j++ {
			l := j - 1
			r := i - j

			total += numTree[l] * numTree[r]
		}

		numTree[i] = total
	}

	return numTree[n]
}
