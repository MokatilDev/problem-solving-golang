package minimum_size_subarray_sum

import "math"

func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt
	sum := 0
	l := 0

	for r := range nums {
		sum += nums[r]

		for target <= sum {
			if res > r-l+1 {
				res = r - l + 1
			}

			sum -= nums[l]
			l++
		}
	}

	if res == math.MaxInt {
		res = 0
	}

	return res
}
