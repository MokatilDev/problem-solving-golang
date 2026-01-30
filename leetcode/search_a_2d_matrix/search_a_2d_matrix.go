package search_a_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	arr := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		if matrix[i][len(matrix[i])-1] < target {
			continue
		}

		arr = append(arr, matrix[i]...)
	}

	l, r := 0, len(arr)-1

	for l <= r {
		mid := (l + r) / 2
		if arr[mid] < target {
			l = mid + 1
		} else if target < arr[mid] {
			r = mid - 1
		} else {
			return true
		}
	}

	return false
}
