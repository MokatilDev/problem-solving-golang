package combinations

func combine(n int, k int) [][]int {
	var res [][]int
	var backtracking func(start int, path []int)
	backtracking = func(start int, path []int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, path)
			return
		}

		for i := start; i <= n; i++ {
			path = append(path, i)
			backtracking(i+1, path)
			path = path[:len(path)-1]
		}
	}

	backtracking(1, []int{})

	return res
}
