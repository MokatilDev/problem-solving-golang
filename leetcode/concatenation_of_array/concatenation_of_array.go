package concatenation_of_array

func getConcatenation(nums []int) []int {
	n := len(nums)
	res := make([]int, 2*n)

	for i, num := range nums {
		res[i], res[i+n] = num, num
	}

	return res
}
