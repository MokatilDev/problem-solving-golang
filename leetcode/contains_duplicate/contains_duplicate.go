package contains_duplicate

func containsDuplicate(nums []int) bool {
	var res = false
	m := make(map[int]bool)

	for _, num := range nums {
		if m[num] == true {
			res = true
			break
		}

		m[num] = true
	}

	return res
}
