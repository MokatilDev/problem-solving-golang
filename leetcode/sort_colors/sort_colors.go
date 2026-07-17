package sort_colors

func sortColors(nums []int) {
	zeros, ones := 0, 0

	for i := range nums {
		if nums[i] == 0 {
			zeros++
		}
		if nums[i] == 1 {
			ones++
		}
	}

	for i := 0; i < zeros; i++ {
		nums[i] = 0
	}

	for i := zeros; i < zeros+ones; i++ {
		nums[i] = 1
	}

	for i := zeros + ones; i < len(nums); i++ {
		nums[i] = 2
	}
}
