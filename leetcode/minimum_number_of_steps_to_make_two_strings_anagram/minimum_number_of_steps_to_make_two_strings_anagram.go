package minimum_number_of_steps_to_make_two_strings_anagram

func minSteps(s string, t string) int {
	var res int

	letters := make(map[rune]int)

	for _, c := range s {
		letters[c]++
	}

	for _, c := range t {
		if letters[c] == 0 {
			res++
			continue
		}

		letters[c]--
	}

	return res
}
