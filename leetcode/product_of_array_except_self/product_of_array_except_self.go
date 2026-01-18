package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	r := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= r
		r *= nums[i]
	}

	return res
}
