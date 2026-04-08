package range_sum_of_sorted_subarray_sums

import (
	"slices"
)

func rangeSum(nums []int, n int, left int, right int) int {
	sums := make([]int, 0, (n*(n+1))/2)
	res := 0

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			sums = append(sums, sum)
		}
	}

	slices.Sort(sums)

	for i := left - 1; i < right; i++ {
		res += sums[i]
	}

	return res % 1_000_000_007
}
