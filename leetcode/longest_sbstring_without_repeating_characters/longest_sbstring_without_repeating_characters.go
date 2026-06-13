package longest_sbstring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	var res int
	set := make(map[byte]int)
	left := 0

	for right := 0; right < len(s); right++ {
		char := s[right]

		if index, exists := set[char]; exists && index >= left {
			left = index + 1
		}

		set[char] = right

		currLength := right - left + 1
		res = max(res, currLength)
	}

	return res
}
