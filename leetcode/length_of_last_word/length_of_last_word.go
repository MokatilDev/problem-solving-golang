package length_of_last_word

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	var res int

	arr := strings.Split(s, " ")
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == "" {
			continue
		} else {
			res = len(arr[i])
			break
		}
	}

	return res
}
