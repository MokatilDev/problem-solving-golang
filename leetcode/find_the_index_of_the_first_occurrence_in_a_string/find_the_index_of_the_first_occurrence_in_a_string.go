package find_the_index_of_the_first_occurrence_in_a_string

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		if i+len(needle) > len(haystack) {
			break
		}

		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
