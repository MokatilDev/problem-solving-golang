package valid_anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var res = true
	letters := make(map[rune]int)

	for _, c := range s {
		letters[c]++
	}

	for _, c := range t {
		letters[c]--

		if letters[c] < 0 {
			res = false
			break
		}
	}

	return res
}
