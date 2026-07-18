package maximum_value_of_an_alternating_sequence

func maximumValue(n int, s int, m int) int64 {
	if n == 1 {
		return int64(s)
	}

	if m == 1 {
		return int64(s + 1)
	}

	last := n - 1

	if last%2 == 0 {
		last--
	}

	a := last / 2

	return int64(s + a*(m-1) + m)
}
