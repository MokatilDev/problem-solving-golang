package find_missing_elements

func findMissingElements(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	var res []int
	hashMap := make(map[int]bool)

	maxNum, minNum := nums[0], nums[0]
	for _, n := range nums {
		hashMap[n] = true

		if n <= minNum {
			minNum = n
		}

		if maxNum <= n {
			maxNum = n
		}
	}

	for n := minNum; n <= maxNum; n++ {
		if !hashMap[n] {
			res = append(res, n)
		}
	}

	return res
}
