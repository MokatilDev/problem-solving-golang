package pow_x_n

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	var calcule func(x float64, n int) float64

	calcule = func(x float64, n int) float64 {
		if n == 0 {
			return 1.0
		}

		half := calcule(x, n/2)
		if n%2 == 0 {
			return half * half
		}

		return half * half * x
	}

	res := calcule(x, n)

	if n < 0 {
		res = 1 / res
	}

	return res
}
