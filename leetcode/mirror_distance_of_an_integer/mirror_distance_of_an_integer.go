package mirror_distance_of_an_integer

import (
	"math"
	"strconv"
)

// Methode 1
func mirrorDistance(n int) int {
	var res int

	s := strconv.Itoa(n)
	reverse := ""

	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}

	r, _ := strconv.Atoi(reverse)
	res = int(math.Abs(float64(n - int(r))))

	return res
}

// Methode 2
func mirrorDistance2(n int) int {
	mirror, temp := 0, n

	for temp > 0 {
		mirror = mirror*10 + temp%10
		temp = temp / 10
	}

	if n < mirror {
		return mirror - n
	}

	return n - mirror
}
