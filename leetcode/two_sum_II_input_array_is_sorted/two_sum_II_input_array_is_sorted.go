package two_sum_II_input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	res := make([]int, 0, 2)

	set := make(map[int]int)

	for i, num := range numbers {
		set[num] = i
	}

	for i, num := range numbers {
		if v, ok := set[target-num]; ok {
			res = append(res, []int{i + 1, v + 1}...)
			break
		}
	}

	return res
}
