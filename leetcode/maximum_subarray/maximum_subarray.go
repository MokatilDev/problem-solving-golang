package maximum_subarray

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	res := nums[0]
	currSum := 0

	for i := range nums {
		if currSum < 0 {
			currSum = 0
		}

		currSum += nums[i]
		res = max(res, currSum)
	}

	return res
}
