package transpose_matrix

func transpose(matrix [][]int) [][]int {
	n := len(matrix)
	p := len(matrix[0])

	res := make([][]int, 0, p)

	for j := range p {
		arr := []int{}
		for i := range n {
			arr = append(arr, matrix[i][j])
		}
		res = append(res, arr)
	}

	return res
}
