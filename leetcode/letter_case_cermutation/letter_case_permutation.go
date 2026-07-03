package letter_case_cermutation

func letterCasePermutation(s string) []string {
	var res []string

	var backtracking func(index int, str []rune)
	backtracking = func(index int, str []rune) {
		if index == len(s) {
			res = append(res, string(str))
			return
		}

		current := rune(s[index])

		if current >= '0' && current <= '9' {
			str = append(str, current)
			backtracking(index+1, str)

			return
		}

		str = append(str, current)
		backtracking(index+1, str)

		str = str[:len(str)-1]

		str = append(str, current^32)
		backtracking(index+1, str)
	}

	backtracking(0, []rune{})

	return res
}
