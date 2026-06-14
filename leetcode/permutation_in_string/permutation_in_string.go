package permutation_in_string

// Method 1 : Hash Map
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	set1 := make(map[byte]int)
	set2 := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		set1[s1[i]]++
		set2[s2[i]]++
	}

	if match(set1, set2) {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		set2[s2[i]]++
		set2[s2[i-len(s1)]]--

		if set2[s2[i-len(s1)]] == 0 {
			delete(set2, s2[i-len(s1)])
		}

		if match(set1, set2) {
			return true
		}
	}

	return false
}

func match(a, b map[byte]int) bool {
	if len(a) != len(b) {
		return false
	}

	for key, value := range a {
		if b[key] != value {
			return false
		}
	}

	return true
}

// Method 2 : Array
/* func checkInclusion(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	arr1 := make([]int, 26)
	arr2 := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		arr1[s1[i]-'a']++
		arr1[s2[i]-'a']++
	}

	if matchArr(arr1, arr2) {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		arr1[s1[i]-'a']++
		arr2[s2[i-len(s1)]-'a']--

		if matchArr(arr1, arr2) {
			return true
		}
	}

	return false
}


func matchArr(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range len(a) {
		if a[i] != b[i] {
			return false
		}
	}

	return true
} */
