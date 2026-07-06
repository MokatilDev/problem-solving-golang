package remove_covered_intervals

import (
	"math"
	"sort"
)

func removeCoveredIntervals(intervals [][]int) int {
	if len(intervals) == 0 || len(intervals[0]) != 2 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	res := 0
	maximum := math.MinInt32

	for i := range intervals {
		if intervals[i][1] > maximum {
			res++
			maximum = intervals[i][1]
		}
	}

	return res
}
