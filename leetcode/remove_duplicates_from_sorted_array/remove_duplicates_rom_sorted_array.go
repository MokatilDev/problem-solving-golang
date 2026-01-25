package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	i := 0

	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
