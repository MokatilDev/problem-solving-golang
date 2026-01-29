package sort_an_array

func sortArray(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(values []int) []int {
	if len(values) < 2 {
		return values
	}

	first := mergeSort(values[:(len(values)-1)/2])
	second := mergeSort(values[(len(values)-1)/2:])

	return merge(first, second)
}

func merge(first, second []int) []int {
	res := make([]int, 0, len(first)+len(second))

	i := 0
	j := 0

	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			res = append(res, first[i])
			i++
		} else {
			res = append(res, second[j])
			j++
		}
	}

	for ; i < len(first); i++ {
		res = append(res, first[i])
	}

	for ; j < len(second); j++ {
		res = append(res, second[j])
	}

	return res
}
