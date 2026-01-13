package integer_to_roman

import "strings"

func intToRoman(num int) string {
	var res strings.Builder

	r := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}

	v := []int{1000, 100, 10, 1}

	for _, d := range v {
		if num%d >= 0 {
			if num/d == 9 {
				res.WriteString(r[d] + r[d*10])
			} else if num/d == 4 {
				res.WriteString(r[d] + r[d*5])
			} else if num/d >= 5 {
				res.WriteString(r[d*5])
				res.WriteString(strings.Repeat(r[d], num/d-5))
			} else {
				res.WriteString(strings.Repeat(r[d], num/d))
			}
		}

		num %= d
	}

	return res.String()
}
