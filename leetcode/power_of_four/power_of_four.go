package power_of_four

func isPowerOfFour(n int) bool {
	if n < 0 {
		return false
	}

	if n == 1 {
		return true
	}

	for n%4 == 0 {
		n = n / 4
	}

	return n == 1
}
