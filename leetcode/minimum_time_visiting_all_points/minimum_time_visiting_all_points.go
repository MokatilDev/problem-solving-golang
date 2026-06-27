package minimum_time_visiting_all_points

import "math"

func minTimeToVisitAllPoints(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	var res int

	x1, y1 := points[0][0], points[0][1]

	for len(points) >= 1 {
		points = points[1:]
		x2, y2 := points[0][0], points[0][1]
		res += int(max(math.Abs(float64(y2-y1)), math.Abs(float64(x2-x1))))
		x1, y1 = x2, y2
	}

	return res
}
