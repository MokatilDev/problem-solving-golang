package find_minimum_in_rotated_sorted_array

import "math"

func findMin(nums []int) int {
	res := math.MaxInt32
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] >= nums[r] {
			l = mid + 1
		} else {
			r = mid - 1
		}

		if res > nums[mid] {
			res = nums[mid]
		}
	}

	return res
}
