package sqrt

func mySqrt(x int) int {
	ans := -1
	l, r := 0, x

	for l <= r {
		mid := l + (r-l)/2
		midSqrt := mid * mid

		if midSqrt == x {
			ans = mid
			break
		} else if midSqrt > x {
			r = mid - 1
		} else {
			ans = mid
			l = mid + 1
		}
	}

	return ans
}
