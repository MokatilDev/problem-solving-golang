package minimum_distance_to_the_target_element

import "math"

func getMinDistance(nums []int, target int, start int) int {
	res := len(nums) - 1

	for i := range nums {
		if nums[i] == target {
			v := i - start
			abs := int(math.Abs(float64(v)))

			if abs < res {
				res = abs
				continue
			}
		}
	}

	return res
}
