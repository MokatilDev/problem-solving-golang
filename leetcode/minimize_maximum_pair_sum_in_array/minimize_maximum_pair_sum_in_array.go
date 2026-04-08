package minimize_maximum_pair_sum_in_array

import (
	"slices"
)

func minPairSum(nums []int) int {
	slices.Sort(nums)

	maximum := 0
	for i := 0; i < len(nums)/2; i++ {
		sum := nums[i] + nums[len(nums)-i-1]
		if maximum < sum {
			maximum = sum
		}
	}

	return maximum
}
