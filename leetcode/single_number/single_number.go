package single_number

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}

	return 0 ^ res
}
