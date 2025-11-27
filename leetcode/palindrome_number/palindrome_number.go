package palindrome

import (
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var str string = strconv.Itoa(x)
	var rev []string
	isCorrect := false

	for i := len(str) - 1; i >= 0; i-- {
		rev = append(rev, strings.Split(str, "")[i])
	}

	if strings.Join(rev, "") == str {
		isCorrect = true
	}

	return isCorrect
}
