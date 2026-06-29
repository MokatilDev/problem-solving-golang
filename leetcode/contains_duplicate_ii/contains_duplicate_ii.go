package contains_duplicate_ii

func containsNearbyDuplicate(nums []int, k int) bool {
	hashMap := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := hashMap[nums[i]]; ok {
			return true
		} else {
			hashMap[nums[i]] = true
		}

		if len(hashMap) > k {
			delete(hashMap, nums[i-k])
		}
	}

	return false
}
