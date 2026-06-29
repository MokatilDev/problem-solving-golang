package longest_mountain_in_array

func longestMountain(arr []int) int {
	var res int

	for i := 1; i < len(arr)-1; i++ {
		var left, right int

		if arr[i-1] < arr[i] && arr[i+1] < arr[i] {
			left, right = i, i

			for left > 0 && arr[left] > arr[left-1] {
				left--
			}

			for right < len(arr)-1 && arr[right] > arr[right+1] {
				right++
			}

			res = max(res, right-left+1)
		}
	}

	return res
}
