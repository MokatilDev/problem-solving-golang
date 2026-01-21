package pascals_triangle

func generate(numRows int) [][]int {
	res := [][]int{{1}}

	for i := range numRows {
		var temp []int
		arr := res[i]
		for j := range len(res[i]) + 1 {
			if j == 0 || j == len(res[i]) {
				temp = append(temp, 1)
			} else {
				temp = append(temp, arr[j]+arr[j-1])
			}
		}
		res = append(res, temp)
	}

	return res
}
