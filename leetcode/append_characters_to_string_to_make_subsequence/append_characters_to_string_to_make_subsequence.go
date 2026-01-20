package append_characters_to_string_to_make_subsequence

func appendCharacters(s string, t string) int {
	var res int

	i := 0
	j := 0

	for i < len(s) || j < len(t) {
		if s[i] == t[j] {
			j++
			i++
		} else {
			break
		}
	}

	res = len(t) - j

	return res
}
