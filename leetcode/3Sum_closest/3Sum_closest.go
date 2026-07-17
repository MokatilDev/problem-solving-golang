package three_sum_closest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return -1
	}

	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-1; i++ {
		l, r := i+1, len(nums)-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return sum
			}

			if math.Abs(float64(target)-float64(sum)) < math.Abs(float64(target)-float64(res)) {
				res = sum
			}

			if sum < target {
				l++
			} else {
				r--
			}
		}
	}

	return res
}
