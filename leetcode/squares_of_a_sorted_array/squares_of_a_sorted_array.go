package squares_of_a_sorted_array

func sortedSquares(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	n := len(nums)
	res := make([]int, n)

	left, right := 0, n-1

	for i := n - 1; i >= 0; i-- {
		rightNum := nums[right] * nums[right]
		leftNum := nums[left] * nums[left]

		if leftNum > rightNum {
			res[i] = leftNum
			left++
		} else {
			res[i] = rightNum
			right--
		}
	}

	return res
}
