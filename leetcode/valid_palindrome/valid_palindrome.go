package valid_palindrome

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")

	for _, char := range s {
		if !unicode.IsDigit(char) && !unicode.IsLetter(char) {
			s = strings.ReplaceAll(s, string(char), "")
		}
	}

	for i := 0; i < int(len(s)/2); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
