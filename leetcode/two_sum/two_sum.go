package two_sum

func twoSum(nums []int, target int) []int {
	var res []int

	values := make(map[int]int)

	for i, num := range nums {
		values[num] = i
	}

	for i, num := range nums {
		if index, ok := values[target-num]; ok {
			if index != i {
				res = []int{index, i}
				break
			}
			continue
		}
	}

	return res
}
