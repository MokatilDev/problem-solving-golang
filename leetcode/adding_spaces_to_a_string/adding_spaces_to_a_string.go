package adding_spaces_to_a_string

func addSpaces(s string, spaces []int) string {
	n, m := len(s), len(spaces)
	res := make([]byte, n+m)

	space_index := 0
	j := 0

	for i := range n {
		if space_index < m && i == spaces[space_index] {
			res[j] = ' '

			space_index++
			j++
		}

		res[j] = s[i]
		j++
	}

	return string(res)
}
