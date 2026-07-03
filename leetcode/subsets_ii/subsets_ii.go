package subsets_ii

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	var backtarking func(start int, path []int)

	backtarking = func(start int, path []int) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)

		for i := start; i < len(nums); i++ {
			if i > start && nums[i-1] == nums[i] {
				continue
			}

			path = append(path, nums[i])
			backtarking(i+1, path)
			path = path[:len(path)-1]
		}
	}

	sort.Ints(nums)
	backtarking(0, []int{})

	return res
}
