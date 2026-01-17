package decode_string

import (
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	res := []string{}

	for _, char := range s {
		if string(char) != "]" {
			res = append(res, string(char))
		} else {
			str := ""
			for len(res) > 0 && res[len(res)-1] != "[" {
				str = res[len(res)-1] + str
				res = res[:len(res)-1]
			}

			if len(res) > 0 {
				res = res[:len(res)-1]
			}

			var k_str string
			for len(res) > 0 && isDigit(res[len(res)-1]) {
				k_str = res[len(res)-1] + k_str
				res = res[:len(res)-1]
			}

			k, _ := strconv.Atoi(k_str)
			repeated := strings.Repeat(str, k)
			res = append(res, repeated)
		}
	}

	return strings.Join(res, "")
}

func isDigit(s string) bool {
	if len(s) == 0 {
		return false
	}
	return unicode.IsDigit(rune(s[0]))
}
