package subsets

func subsets(nums []int) [][]int {
	res := [][]int{}
	subset := []int{}

	var subsetRecur func(i int)

	subsetRecur = func(i int) {
		if i == len(nums) {
			temp := make([]int, len(subset))
			copy(temp, subset)

			res = append(res, temp)
			return
		}

		subset = append(subset, nums[i])
		subsetRecur(i + 1)

		subset = subset[:len(subset)-1]
		subsetRecur(i + 1)
	}

	subsetRecur(0)

	return res
}
