package replace_elements_with_greatest_element_on_right_side

func replaceElements(arr []int) []int {
	l := len(arr)

	if l == 1 {
		return []int{-1}
	}

	res := make([]int, l)

	for i := 0; i < l-1; i++ {
		max := arr[i+1]

		for j := i + 2; j < l; j++ {
			if max < arr[j] {
				max = arr[j]
			}
		}

		res[i] = max
	}

	res[l-1] = -1

	return res
}
