package score_of_a_string

import (
	"math"
)

func scoreOfString(s string) int {
	var res int

	for i := 0; i < len(s)-1; i++ {
		res += int(math.Abs(float64(int(s[i]) - int(s[i+1]))))
	}

	return res
}
