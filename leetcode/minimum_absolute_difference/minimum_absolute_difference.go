package minimum_absolute_difference

import (
	"math"
	"slices"
)

func minimumAbsDifference(arr []int) [][]int {
	if len(arr) < 2 {
		return nil
	}

	slices.Sort(arr)

	var res [][]int
	minDiff := math.MaxInt

	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]

		if diff < minDiff {
			minDiff = diff
			res = [][]int{{arr[i], arr[i+1]}}
		} else if diff == minDiff {
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}

	return res
}
