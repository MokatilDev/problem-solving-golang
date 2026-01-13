package roman_to_integer

func romanToInt(s string) int {
	var res int

	romans := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	for i := 0; i < len(s)-1; i++ {
		a := romans[string(s[i])]
		b := romans[string(s[i+1])]

		if a < b {
			res += -a

		} else {
			res += a
		}

	}

	return res + romans[string(s[len(s)-1])]
}
