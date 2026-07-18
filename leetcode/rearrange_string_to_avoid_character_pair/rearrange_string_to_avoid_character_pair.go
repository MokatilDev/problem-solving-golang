package rearrange_string_to_avoid_character_pair

func rearrangeString(s string, x byte, y byte) string {
	var t string

	for i := range s {
		char := s[i]
		if char == y {
			t = string(char) + t
			continue
		}
		t += string(char)
	}

	return t
}
