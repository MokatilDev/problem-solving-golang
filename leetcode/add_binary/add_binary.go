package addBinary

import (
	"strings"
)

const ZERO, ONE = '0', '1'

func addition(x, y, inCarry byte) (sum, outCarry byte) {
	if x == ZERO && y == ZERO {
		sum = ZERO
		outCarry = ZERO
	} else if x == ONE && y == ONE {
		sum = ZERO
		outCarry = ONE
	} else {
		sum = ONE
		outCarry = ZERO
	}

	if inCarry == ONE {
		if sum == ONE {
			sum = ZERO
			outCarry = ONE
		} else {
			sum = ONE
		}
	}

	return sum, outCarry
}

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1

	carry := byte(ZERO)
	var s strings.Builder
	s.Grow(max(i, j) + 1)

	for i >= 0 || j >= 0 || carry == ONE {
		var x, y byte

		if i >= 0 {
			x = a[i]
		} else {
			x = ZERO
		}

		if j >= 0 {
			y = b[j]
		} else {
			y = ZERO
		}
		var sum byte
		sum, carry = addition(x, y, carry)

		s.WriteByte(sum)
		i--
		j--
	}

	res := s.String()
	runnes := []rune(res)

	for i, j := 0, len(runnes)-1; i < j; i, j = i+1, j-1 {
		runnes[i], runnes[j] = runnes[j], runnes[i]
	}

	return string(runnes)
}
