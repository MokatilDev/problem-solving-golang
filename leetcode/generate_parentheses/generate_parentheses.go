package generate_parentheses

func generateParenthesis(n int) []string {
	var res []string
	backtrack(&res, "", 0, 0, n)
	return res
}

func backtrack(res *[]string, current string, open, close, n int) {
	if len(current) == n*2 {
		*res = append(*res, current)
		return
	}

	if open < n {
		backtrack(res, current+"(", open+1, close, n)
	}

	if close < open {
		backtrack(res, current+")", open, close+1, n)
	}
}
