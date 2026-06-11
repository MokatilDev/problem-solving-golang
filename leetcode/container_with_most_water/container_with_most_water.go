package container_with_most_water

func maxArea(height []int) int {
	var res int
	var area int

	l, r := 0, len(height)-1

	for l < r {
		area = (r - l) * min(height[r], height[l])
		res = max(res, area)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}
