package count_good_numbers

const mod = 1_000_000_007

func countGoodNumbers(n int64) int {
	even := (n + 1) / 2
	odd := n / 2

	var calcule func(x, n int64) int64
	calcule = func(x, n int64) int64 {
		if n == 0 {
			return 1
		}
		half := calcule(x, n/2)
		result := (half * half) % mod
		if n%2 != 0 {
			result = (result * x) % mod
		}
		return result
	}

	res := (calcule(5, even) * calcule(4, odd)) % mod
	return int(res)
}
