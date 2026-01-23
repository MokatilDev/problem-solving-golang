package monotonic_array

func isMonotonic(nums []int) bool {
	increasing, decreasing := false, false

	for i := 0; i < len(nums)-1; i++ {

		if nums[i] < nums[i+1] {
			increasing = true
		}
		if nums[i] > nums[i+1] {
			decreasing = true
		}

	}

	return !(increasing && decreasing)
}
