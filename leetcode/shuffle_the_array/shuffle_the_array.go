package shuffle_the_array

func shuffle(nums []int, n int) []int {
	res := make([]int, 2*n)

	for i := range n {
		res[2*i] = nums[i]
		res[2*i+1] = nums[i+n]
	}

	return res
}
