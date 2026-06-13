package longest_repeating_character_replacement

func characterReplacement(s string, k int) int {
	var maxLen int
	set := make(map[byte]int)

	left := 0
	maxFreq := 0

	for right := 0; right < len(s); right++ {
		char := s[right]
		set[char]++

		if set[char] > maxFreq {
			maxFreq = set[char]
		}

		if (right-left+1)-maxFreq > k {
			set[s[left]]--
			left++
		}

		if (right - left + 1) > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}
