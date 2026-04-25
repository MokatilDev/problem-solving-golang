package merge_strings_alternately

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	var res strings.Builder

	i := 0
	j := 0

	for i < len(word1) && j < len(word2) {
		res.WriteByte(word1[i])
		res.WriteByte(word2[j])

		i++
		j++
	}

	if i != len(word1) {
		for k := i; k < len(word1); k++ {
			res.WriteByte(word1[k])
		}
	}

	if j != len(word2) {
		for k := j; k < len(word2); k++ {
			res.WriteByte(word2[k])
		}
	}

	return res.String()
}
