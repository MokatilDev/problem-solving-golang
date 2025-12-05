package remove_element

func removeElement(nums []int, val int) int {
	k := 0

	for _, num := range nums {
		if num != val {
			nums[k] = val
			k++
			continue
		}
	}

	return k
}
