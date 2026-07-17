package four_sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			sum1 := target - nums[i] - nums[j]
			l, r := j+1, len(nums)-1

			for l < r {
				sum2 := nums[l] + nums[r]
				if sum2 == sum1 {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--

					for l < r && nums[l] == nums[l-1] {
						l++
					}

					for l < r && nums[r] == nums[r+1] {
						r--
					}
				} else if sum2 > sum1 {
					r--
				} else {
					l++
				}
			}
		}
	}

	return res
}
