package maximum_product_subarray

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res, maxProd, minProd := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxProd, minProd = minProd, maxProd
		}

		maxProd = max(nums[i], maxProd*nums[i])
		minProd = min(nums[i], minProd*nums[i])

		res = max(res, maxProd)
	}

	return res
}
