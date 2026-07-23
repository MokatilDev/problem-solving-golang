package zigzag_conversion

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var ans strings.Builder
	chars := make([][]byte, numRows)

	current := 0
	direction := -1

	for i := range s {
		chars[current] = append(chars[current], s[i])
		if current == 0 || current == numRows-1 {
			direction = -direction
		}

		current += direction
	}

	for i := range chars {
		for _, char := range chars[i] {
			ans.WriteByte(char)
		}
	}

	return ans.String()
}
