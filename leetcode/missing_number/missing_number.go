package missing_number

// O(nlog(n))
/* func missingNumber(nums []int) int {
	slices.Sort(nums)

	for i, v := range nums {
		if i != v {
			return v - 1
		}

		if v == len(nums)-1 {
			return v + 1
		}
	}

	return -1
} */

func missingNumber(nums []int) int {
	n := len(nums)
	expectedSum := (n * (n + 1)) / 2

	actualSum := 0
	for _, v := range nums {
		actualSum += v
	}

	return expectedSum - actualSum
}
