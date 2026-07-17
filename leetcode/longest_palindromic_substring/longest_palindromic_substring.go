package longest_palindromic_substring

func longestPalindrome(s string) string {
	var res string

	for i := 0; i < len(s); i++ {
		s1 := isPalindromic(s, i, i)
		s2 := isPalindromic(s, i, i+1)

		if len(res) < len(s1) {
			res = s1
		}

		if len(res) < len(s2) {
			res = s2
		}
	}

	return string(res)
}

func isPalindromic(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return s[l+1 : r]
}
